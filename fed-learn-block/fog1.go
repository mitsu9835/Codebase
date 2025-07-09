package main

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"

    bls "github.com/kilic/bls12-381"
    "fed-learn-block/functions"
)

func main() {
    const workspace = "tmp/fog1"
    deviceFile := filepath.Join(workspace, "device_output.json")
    fogFile := filepath.Join(workspace, "fog1_output.json")
    pubFile := filepath.Join(workspace, "fog1_pubkey.json")

    // 1) Clean up any old state & recreate workspace
    if err := os.RemoveAll(workspace); err != nil {
        log.Fatalf("cleanup %s: %v", workspace, err)
    }
    if err := os.MkdirAll(workspace, 0755); err != nil {
        log.Fatalf("mkdir %s: %v", workspace, err)
    }

    // 2) Prepare BLS G1 context once
    g1 := bls.NewG1()

    for {
        // 3) Wait for a device to drop its JSON
        fmt.Printf("FOG1: waiting for %s …\n", deviceFile)
        for {
            if _, err := os.Stat(deviceFile); err == nil {
                break
            }
            time.Sleep(500 * time.Millisecond)
        }
        fmt.Printf("FOG1: detected %s, loading…\n", deviceFile)

        // 4) Load and parse device_output.json
        data, err := ioutil.ReadFile(deviceFile)
        if err != nil {
            log.Fatalf("read %s: %v", deviceFile, err)
        }
        var in struct {
            CommitG string `json:"commit_g"`
            CommitN string `json:"commit_n"`
            CommitM string `json:"commit_m"`
        }
        if err := json.Unmarshal(data, &in); err != nil {
            log.Fatalf("unmarshal %s: %v", deviceFile, err)
        }

        // 5) Decode & decompress the commitments
        cgBytes, _ := hex.DecodeString(in.CommitG)
        commitG, err := g1.FromCompressed(cgBytes)
        if err != nil {
            log.Fatalf("decompress commit_g: %v", err)
        }
        cnBytes, _ := hex.DecodeString(in.CommitN)
        commitN, err := g1.FromCompressed(cnBytes)
        if err != nil {
            log.Fatalf("decompress commit_n: %v", err)
        }
        cmBytes, _ := hex.DecodeString(in.CommitM)
        commitM, err := g1.FromCompressed(cmBytes)
        if err != nil {
            log.Fatalf("decompress commit_m: %v", err)
        }

        // 6) Verify commitM == commitG + commitN
        sum := g1.New()
        g1.Add(sum, commitG, commitN)
        if !g1.Equal(sum, commitM) {
            log.Fatal("FOG1: Commitment verification FAILED")
        }
        fmt.Println("FOG1: Commitment verification succeeded")

        // 7) Generate an ECDSA keypair for Fog and publish the public key
        priv, pub := functions.GenerateKeyPair()
        pubOut := struct {
            X string `json:"X"`
            Y string `json:"Y"`
        }{
            X: fmt.Sprintf("%x", pub.X.Bytes()),
            Y: fmt.Sprintf("%x", pub.Y.Bytes()),
        }
        pubData, _ := json.MarshalIndent(pubOut, "", "  ")
        if err := ioutil.WriteFile(pubFile, pubData, 0644); err != nil {
            log.Fatalf("write %s: %v", pubFile, err)
        }

        // 8) Sign the compressed commitG and write fog1_output.json
        sig, err := functions.SignCommitmentECDSA(priv, cgBytes)
        if err != nil {
            log.Fatalf("sign commit_g: %v", err)
        }
        outSig := struct {
            R string `json:"R"`
            S string `json:"S"`
        }{
            R: fmt.Sprintf("%x", sig.R),
            S: fmt.Sprintf("%x", sig.S),
        }
        sigData, _ := json.MarshalIndent(outSig, "", "  ")
        if err := ioutil.WriteFile(fogFile, sigData, 0644); err != nil {
            log.Fatalf("write %s: %v", fogFile, err)
        }
        fmt.Printf("FOG1: Wrote %s and %s\n", fogFile, pubFile)

        // 9) Remove the device file so that next run will wait again
        if err := os.Remove(deviceFile); err != nil {
            log.Fatalf("remove %s: %v", deviceFile, err)
        }
        // loop back for the next device (A or B)
    }
}
