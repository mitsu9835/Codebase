package main

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"

    bls "github.com/kilic/bls12-381"
    "fed-learn-block/functions"
)

func main() {
    // 1) Read device’s JSON output
    data, err := ioutil.ReadFile("tmp/device_output.json")
    if err != nil {
        log.Fatalf("read device_output.json: %v", err)
    }
    var in struct {
        CommitG string `json:"commit_g"`
        CommitN string `json:"commit_n"`
        CommitM string `json:"commit_m"`
    }
    if err := json.Unmarshal(data, &in); err != nil {
        log.Fatalf("unmarshal device_output.json: %v", err)
    }

    // 2) Decode and decompress commit_g, commit_n, commit_m
    g1 := bls.NewG1()

    cgBytes, err := hex.DecodeString(in.CommitG)
    if err != nil {
        log.Fatalf("decode commit_g: %v", err)
    }
    commitG, err := g1.FromCompressed(cgBytes)
    if err != nil {
        log.Fatalf("decompress commit_g: %v", err)
    }

    cnBytes, err := hex.DecodeString(in.CommitN)
    if err != nil {
        log.Fatalf("decode commit_n: %v", err)
    }
    commitN, err := g1.FromCompressed(cnBytes)
    if err != nil {
        log.Fatalf("decompress commit_n: %v", err)
    }

    cmBytes, err := hex.DecodeString(in.CommitM)
    if err != nil {
        log.Fatalf("decode commit_m: %v", err)
    }
    commitM, err := g1.FromCompressed(cmBytes)
    if err != nil {
        log.Fatalf("decompress commit_m: %v", err)
    }

    // 3) Verify commitM == commitG + commitN
    sum := g1.New()
    g1.Add(sum, commitG, commitN)
    if g1.Equal(sum, commitM) {
        fmt.Println("Commitment verification succeeded")
    } else {
        log.Fatal("Commitment verification FAILED")
    }

    // 4) Generate an ECDSA keypair for Fog
    priv, pub := functions.GenerateKeyPair()

    // 5) Persist Fog’s public key for Device verification
    pubOut := struct {
        X string `json:"X"`
        Y string `json:"Y"`
    }{
        X: fmt.Sprintf("%x", pub.X.Bytes()),
        Y: fmt.Sprintf("%x", pub.Y.Bytes()),
    }
    pubData, err := json.MarshalIndent(pubOut, "", "  ")
    if err != nil {
        log.Fatalf("marshal fog_pubkey.json: %v", err)
    }
    if err := ioutil.WriteFile("tmp/fog_pubkey.json", pubData, 0644); err != nil {
        log.Fatalf("write fog_pubkey.json: %v", err)
    }

    // 6) Sign commitG
    sig, err := functions.SignCommitmentECDSA(priv, cgBytes)
    if err != nil {
        log.Fatalf("sign commit_g: %v", err)
    }

    // 7) Write R and S to fog_output.json
    outSig := struct {
        R string `json:"R"`
        S string `json:"S"`
    }{
        R: fmt.Sprintf("%x", sig.R),
        S: fmt.Sprintf("%x", sig.S),
    }
    sigData, err := json.MarshalIndent(outSig, "", "  ")
    if err != nil {
        log.Fatalf("marshal fog_output.json: %v", err)
    }
    if err := ioutil.WriteFile("tmp/fog_output.json", sigData, 0644); err != nil {
        log.Fatalf("write fog_output.json: %v", err)
    }
    fmt.Println("Wrote tmp/fog_output.json with ECDSA signature")
}
