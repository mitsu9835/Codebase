package main

import (
    "context"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"

    bls "github.com/kilic/bls12-381"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "fed-learn-block/contracts"
    "fed-learn-block/functions"
)

// computeLocalGradient applies ∇ = Σᵢ[(θ·xᵢ − yᵢ) * xᵢ]
func computeLocalGradient(X [][]float64, y []float64, theta []float64) []float64 {
    m := len(X)
    n := len(X[0])
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
    // 1) Load CSV
    f, err := os.Open("train.csv")
    if err != nil {
        log.Fatalf("open train.csv: %v", err)
    }
    defer f.Close()
    reader := csv.NewReader(f)

    // 2) Read header to get feature count
    header, err := reader.Read()
    if err != nil {
        log.Fatalf("read header: %v", err)
    }
    n := len(header) - 1
    fmt.Printf("Detected %d features (excluding target)\n", n)

    // 3) Read all records
    var records [][]string
    for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("read record: %v", err)
        }
        records = append(records, rec)
    }
    m := len(records)
    fmt.Printf("Total examples: %d\n", m)

    // 4) Parse into X, y
    X := make([][]float64, m)
    y := make([]float64, m)
    for i, rec := range records {
        X[i] = make([]float64, n)
        for j := 0; j < n; j++ {
            v, err := strconv.ParseFloat(rec[j], 64)
            if err != nil {
                log.Fatalf("parse feature %q: %v", rec[j], err)
            }
            X[i][j] = v
        }
        tv, err := strconv.ParseFloat(rec[n], 64)
        if err != nil {
            log.Fatalf("parse target %q: %v", rec[n], err)
        }
        y[i] = tv
    }

    // 5) Initialize theta = [0.1, 0.2, …]
    theta := make([]float64, n)
    for i := range theta {
        theta[i] = 0.1 * float64(i+1)
    }
    fmt.Printf("Initialized theta: %v\n", theta)

    // 6) Dot product for first example
    dot := 0.0
    for j := 0; j < n; j++ {
        dot += theta[j] * X[0][j]
    }
    fmt.Printf("Dot product for example #1: %f\n", dot)

    // 7) Compute local gradient
    grad := computeLocalGradient(X, y, theta)
    fmt.Printf("Local gradient: %v\n", grad)

    // --- Ethereum noise generation --- //

    // 8) Read contract address
    addrBytes, err := os.ReadFile("tmp/contractInfo/contract-address")
    if err != nil {
        log.Fatalf("read address: %v", err)
    }
    addr := common.HexToAddress(strings.TrimSpace(string(addrBytes)))
    fmt.Println("Using Store at", addr.Hex())

    // 9) Dial node
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatalf("dial node: %v", err)
    }

    // 10) Bind Store
    store, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatalf("bind contract: %v", err)
    }

    // 11) Fetch one noise sample per feature
    noises := make([]int64, n)
    for j := 0; j < n; j++ {
        nb, err := store.GenerateNoise(&bind.CallOpts{Context: context.Background()})
        if err != nil {
            log.Fatalf("generate_noise[%d]: %v", j, err)
        }
        noises[j] = nb.Int64()
    }
    fmt.Printf("Noise vector η: %v\n", noises)

    // 12) Mask gradient
    masked := functions.MaskGradient(grad, func() []float64 {
        arr := make([]float64, n)
        for i, v := range noises {
            arr[i] = float64(v)
        }
        return arr
    }())
    fmt.Printf("Masked gradient: %v\n", masked)

    // 13) Compute BLS commitments
    g1 := bls.NewG1()
    commitG := functions.CommitVector(grad)
    commitN := functions.CommitVector(func() []float64 {
        arr := make([]float64, n)
        for i, v := range noises {
            arr[i] = float64(v)
        }
        return arr
    }())
    commitM := functions.CommitVector(masked)

    cgBytes := g1.ToCompressed(commitG)
    cnBytes := g1.ToCompressed(commitN)
    cmBytes := g1.ToCompressed(commitM)

    // 14) Write device_output.json
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
        log.Fatalf("marshal JSON: %v", err)
    }
    if err := ioutil.WriteFile("tmp/device_output.json", data, 0644); err != nil {
        log.Fatalf("write device_output.json: %v", err)
    }
    fmt.Println("Wrote tmp/device_output.json")
}
