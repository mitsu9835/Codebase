package main

import (
    "context"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "fed-learn-block/contracts"
)

// computeLocalGradient applies the provided algorithm:
// ∇ = Σᵢ[(θ·xᵢ − yᵢ) * xᵢ]
func computeLocalGradient(X [][]float64, y []float64, theta []float64) []float64 {
    m := len(X)
    n := len(X[0])
    gradient := make([]float64, n)
    for i := 0; i < m; i++ {
        pred := 0.0
        for j := 0; j < n; j++ {
            pred += theta[j] * X[i][j]
        }
        err := pred - y[i]
        for j := 0; j < n; j++ {
            gradient[j] += err * X[i][j]
        }
    }
    return gradient
}

func main() {
    // 1) Open CSV
    f, err := os.Open("train.csv")
    if err != nil {
        log.Fatalf("failed to open train.csv: %v", err)
    }
    defer f.Close()

    reader := csv.NewReader(f)

    // 2) Read header
    header, err := reader.Read()
    if err != nil {
        log.Fatalf("failed to read header: %v", err)
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
            log.Fatalf("failed to read records: %v", err)
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

    // 5) Initialize theta
    theta := make([]float64, n)
    for i := range theta {
        theta[i] = 0.1 * float64(i+1) // example coefficients
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

    // --- Ethereum setup for noise generation --- //

    // 8) Read deployed Store address
    addrBytes, err := os.ReadFile("tmp/contractInfo/contract-address")
    if err != nil {
        log.Fatalf("read address file: %v", err)
    }
    addr := common.HexToAddress(strings.TrimSpace(string(addrBytes)))
    fmt.Println("Using Store at", addr.Hex())

    // 9) Connect to the node
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatalf("dial node: %v", err)
    }

    // 10) Bind the contract
    store, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatalf("bind contract: %v", err)
    }

    // 11) Call generate_noise()
    noiseBig, err := store.GenerateNoise(&bind.CallOpts{Context: context.Background()})
    if err != nil {
        log.Fatalf("generate_noise: %v", err)
    }
    fmt.Printf("Generated noise ηᵢⱼ = %s\n", noiseBig.String())
}
