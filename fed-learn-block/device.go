// device.go
package main

import (
    "context"
    "crypto/ecdsa"
    "crypto/elliptic"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "math/big"
    "os"
    "strconv"
    "strings"
    "time"

    bls "github.com/kilic/bls12-381"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "fed-learn-block/contracts"
    "fed-learn-block/functions"
)

// computeLocalGradient applies ∇ = Σᵢ[(θ·xᵢ − yᵢ) * xᵢ]
func computeLocalGradient(X [][]float64, y []float64, theta []float64) []float64 {
    m, n := len(X), len(X[0])
    grad := make([]float64, n)
    for i := 0; i < m; i++ {
        pred := 0.0
        for j := 0; j < n; j++ {
            pred += theta[j] * X[i][j]
        }
        err := pred - y[i]
        for j := 0; j < n; j++ {
            grad[j] += err * X[i][j]
        }
    }
    return grad
}

func main() {
    // ─── Automatic cleanup of any old Fog files ────────────────────
    os.Remove("tmp/fog_pubkey.json")
    os.Remove("tmp/fog_output.json")

    // 1) Load CSV and parse
    f, err := os.Open("train.csv")
    if err != nil {
        log.Fatalf("open train.csv: %v", err)
    }
    defer f.Close()
    reader := csv.NewReader(f)
    header, _ := reader.Read()
    n := len(header) - 1
    fmt.Printf("Detected %d features (excluding target)\n", n)

    var recs [][]string
    for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("read record: %v", err)
        }
        recs = append(recs, rec)
    }
    m := len(recs)
    fmt.Printf("Total examples: %d\n", m)

    // 2) Parse into X, y slices
    X := make([][]float64, m)
    y := make([]float64, m)
    for i, r := range recs {
        X[i] = make([]float64, n)
        for j := 0; j < n; j++ {
            v, _ := strconv.ParseFloat(r[j], 64)
            X[i][j] = v
        }
        y[i], _ = strconv.ParseFloat(r[n], 64)
    }

    // 3) Initialize theta
    theta := make([]float64, n)
    for i := range theta {
        theta[i] = 0.1 * float64(i+1)
    }
    fmt.Printf("Initialized theta: %v\n", theta)

    // 4) Dot product for first example
    dot := 0.0
    for j := 0; j < n; j++ {
        dot += theta[j] * X[0][j]
    }
    fmt.Printf("Dot product for example #1: %f\n", dot)

    // 5) Compute local gradient
    grad := computeLocalGradient(X, y, theta)
    fmt.Printf("Local gradient: %v\n", grad)

    // ─── Fetch on-chain noise via Store.GenerateNoise() ────────────
    addrBytes, err := os.ReadFile("tmp/contractInfo/contract-address")
    if err != nil {
        log.Fatalf("read address: %v", err)
    }
    addr := common.HexToAddress(strings.TrimSpace(string(addrBytes)))
    fmt.Println("Using Store at", addr.Hex())

    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatalf("dial node: %v", err)
    }
    store, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatalf("bind contract: %v", err)
    }

    noises := make([]float64, n)
    for j := 0; j < n; j++ {
        nb, err := store.GenerateNoise(&bind.CallOpts{Context: context.Background()})
        if err != nil {
            log.Fatalf("generate_noise[%d]: %v", j, err)
        }
        noises[j] = float64(nb.Int64())
    }
    fmt.Printf("Noise vector η: %v\n", noises)

    // ─── Mask, Commit, and Write device_output.json ───────────────
    masked := functions.MaskGradient(grad, noises)
    fmt.Printf("Masked gradient: %v\n", masked)

    g1 := bls.NewG1()
    commitG := functions.CommitVector(grad)
    commitN := functions.CommitVector(noises)
    commitM := functions.CommitVector(masked)

    cgBytes := g1.ToCompressed(commitG)
    cnBytes := g1.ToCompressed(commitN)
    cmBytes := g1.ToCompressed(commitM)

    out := struct {
        MaskedUpdate []float64 `json:"masked_update"`
        CommitG      string    `json:"commit_g"`
        CommitN      string    `json:"commit_n"`
        CommitM      string    `json:"commit_m"`
    }{
        MaskedUpdate: masked,
        CommitG:      fmt.Sprintf("%x", cgBytes),
        CommitN:      fmt.Sprintf("%x", cnBytes),
        CommitM:      fmt.Sprintf("%x", cmBytes),
    }

    data, err := json.MarshalIndent(out, "", "  ")
    if err != nil {
        log.Fatalf("marshal device_output.json: %v", err)
    }
    if err := os.MkdirAll("tmp", 0755); err != nil {
        log.Fatalf("mkdir tmp: %v", err)
    }
    if err := ioutil.WriteFile("tmp/device_output.json", data, 0644); err != nil {
        log.Fatalf("write device_output.json: %v", err)
    }
    fmt.Println("Wrote tmp/device_output.json")

    // ─── Wait for Fog’s output ─────────────────────────────────────
    for {
        if _, err := os.Stat("tmp/fog_pubkey.json"); err == nil {
            break
        }
        time.Sleep(time.Second)
    }
    for {
        if _, err := os.Stat("tmp/fog_output.json"); err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    // ─── Load Fog’s pubkey and signature ───────────────────────────
    pkData, err := ioutil.ReadFile("tmp/fog_pubkey.json")
    if err != nil {
        log.Fatalf("read fog_pubkey.json: %v", err)
    }
    var pk struct{ X, Y string }
    if err := json.Unmarshal(pkData, &pk); err != nil {
        log.Fatalf("unmarshal fog_pubkey.json: %v", err)
    }
    x, _ := new(big.Int).SetString(pk.X, 16)
    yPub, _ := new(big.Int).SetString(pk.Y, 16)
    pubKey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: yPub}

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

    // ─── Verify Fog’s ECDSA signature on CommitG ────────────────
    if functions.VerifySignatureECDSA(pubKey, cgBytes, r, s) {
        fmt.Println("Fog signature valid ✅")
    } else {
        fmt.Println("Fog signature INVALID ❌")
    }
}
