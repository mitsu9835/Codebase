package main

import (
    "context"
    "fmt"
    "log"

    "fed-learn-block/contracts"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // 1. Connect to your node
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatalf("Failed to connect to node: %v", err)
    }
    defer client.Close()

    // 2. Confirm connection
    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatalf("Failed to get chain ID: %v", err)
    }
    fmt.Println("Connected; chain ID:", chainID)

    // 3. Your deployed contract address
    addr := common.HexToAddress("0xC57237Cac8a2A561Ad83bee84691e635738faEdD")

    // 4. Instantiate the binding (use NewContracts)
    contract, err := contracts.NewContracts(addr, client)
    if err != nil {
        log.Fatalf("Failed to instantiate contract: %v", err)
    }

    // 5. Call the GetNonce() method
    noiseBytes, err := contract.GetNonce(nil)
    if err != nil {
        log.Fatalf("Failed to call GetNonce: %v", err)
    }

    // 6. Print the result as hex
    fmt.Printf("Noise (storedNonce): %x\n", noiseBytes)
}
