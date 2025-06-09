package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "fed-learn-block/contracts"
)

func main() {
    // 1) connect
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }

    // 2) load the deployed address
    raw, err := ioutil.ReadFile("./tmp/contractInfo/contract-address")
    if err != nil {
        log.Fatal(err)
    }
    addr := common.HexToAddress(strings.TrimSpace(string(raw)))

    // 3) instantiate
    inst, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatal(err)
    }

    // 4) call
    callOpts := &bind.CallOpts{Context: context.Background()}
    b, err := inst.GetNonce(callOpts)
    if err != nil {
        log.Fatal(err)
    }

    // 5) decode bytes→big.Int
    nonce := new(big.Int).SetBytes(b)
    fmt.Println("Contract nonce is:", nonce.String())
}
