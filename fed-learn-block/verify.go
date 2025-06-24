package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/big"

    "fed-learn-block/functions"
)

func main() {
    // 1) Read the original device_output.json
    data, err := ioutil.ReadFile("tmp/device_output.json")
    if err != nil {
        log.Fatalf("read device_output.json: %v", err)
    }
    var dev struct {
        CommitG string `json:"commit_g"`
    }
    if err := json.Unmarshal(data, &dev); err != nil {
        log.Fatalf("unmarshal device_output.json: %v", err)
    }

    // 2) Decode commit_g
    cgBytes, err := hex.DecodeString(dev.CommitG)
    if err != nil {
        log.Fatalf("decode commit_g: %v", err)
    }

    // 3) Load Fog’s public key
    pkData, err := ioutil.ReadFile("tmp/fog_pubkey.json")
    if err != nil {
        log.Fatalf("read fog_pubkey.json: %v", err)
    }
    var pk struct{ X, Y string }
    if err := json.Unmarshal(pkData, &pk); err != nil {
        log.Fatalf("unmarshal fog_pubkey.json: %v", err)
    }
    x, _ := new(big.Int).SetString(pk.X, 16)
    y, _ := new(big.Int).SetString(pk.Y, 16)
    pubKey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

    // 4) Load Fog’s signature
    sigData, err := ioutil.ReadFile("tmp/fog_output.json")
    if err != nil {
        log.Fatalf("read fog_output.json: %v", err)
    }
    var sig struct{ R, S string }
    if err := json.Unmarshal(sigData, &sig); err != nil {
        log.Fatalf("unmarshal fog_output.json: %v", err)
    }
    r, _ := new(big.Int).SetString(sig.R, 16)
    s, _ := new(big.Int).SetString(sig.S, 16)

    // 5) Verify signature on commit_g
    valid := functions.VerifySignatureECDSA(pubKey, cgBytes, r, s)
    if valid {
        fmt.Println("Fog signature valid ✅")
    } else {
        fmt.Println("Fog signature INVALID ❌")
    }
}
