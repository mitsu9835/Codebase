package main

import (
    "context"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "math/big"
    "os"
    "strconv"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "fed-learn-block/contracts"
)

func main() {
    // 0) Read deployed contract address
    data, err := os.ReadFile("tmp/contractInfo/contract-address")
    if err != nil {
        log.Fatalf("read address file: %v", err)
    }
    addr := common.HexToAddress(strings.TrimSpace(string(data)))
    fmt.Println("Using contract at", addr.Hex())

    // 1) Open train.csv
    f, err := os.Open("train.csv")
    if err != nil {
        log.Fatalf("open train.csv: %v", err)
    }
    defer f.Close()
    reader := csv.NewReader(f)

    // 2) Read header → feature count
    header, err := reader.Read()
    if err != nil {
        log.Fatalf("read header: %v", err)
    }
    n := len(header) - 1

    // 3) Load all rows
    var Xflat []*big.Int
    var y []*big.Int
    for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("read CSV: %v", err)
        }
        // features
        for i := 0; i < n; i++ {
            fval, err := strconv.ParseFloat(rec[i], 64)
            if err != nil {
                log.Fatalf("parse feature %q: %v", rec[i], err)
            }
            iv := int64(fval * 1e6)
            Xflat = append(Xflat, big.NewInt(iv))
        }
        // target
        fval, err := strconv.ParseFloat(rec[n], 64)
        if err != nil {
            log.Fatalf("parse target %q: %v", rec[n], err)
        }
        iv := int64(fval * 1e6)
        y = append(y, big.NewInt(iv))
    }
    m := len(y)

    // 4) Prepare θ = zeros
    theta := make([]*big.Int, n)
    for i := range theta {
        theta[i] = big.NewInt(0)
    }

    // 5) Connect to node
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatalf("dial node: %v", err)
    }

    // 6) Bind contract
    store, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatalf("bind contract: %v", err)
    }

    // 7) Call pure function
    grad, err := store.ComputeGradientCalldata(
        &bind.CallOpts{Context: context.Background()},
        Xflat,
        big.NewInt(int64(m)),
        big.NewInt(int64(n)),
        y,
        theta,
    )
    if err != nil {
        log.Fatalf("ComputeGradientCalldata failed: %v", err)
    }

    // 8) Print gradient
    fmt.Println("Gradient (scale=1e6):")
    for i, v := range grad {
        fmt.Printf("  [%d] = %s\n", i, v.String())
    }
}
