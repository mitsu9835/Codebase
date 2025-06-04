package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"fed-learn-block/contracts"
	"fed-learn-block/functions"
)

var pairing *pbc.Pairing
var g *pbc.Element

func CommContract(paramsByte []byte) {
	// 1. Get the Geth RPC URL and the account JSON
	gethPath, nodeKey := functions.GethPathAndKey()

	// 2. Connect to the private chain
	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
	fmt.Println("connection Returned: ", connection)
	ctx := context.Background()

	// 3. Create a transactor that includes the chain ID (2020)
	chainID := big.NewInt(2020) // your Clique chain ID
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(nodeKey), "Chirag1", chainID)
	functions.CheckError(err)

	// 4. Deploy the Contracts contract
	start := time.Now()
	addr, tx, instContract, err := contracts.DeployContracts(auth, connection)
	functions.CheckError(err)
	fmt.Println("\n addr Returned: ", addr)
	fmt.Println("\n instContract Returned: ", instContract)

	// 5. Save the deployed address for other apps to use
	functions.FileCrWr("./tmp/contractInfo/contract-address", []byte(addr.Hex()))

	// 6. Wait for the deployment transaction to be mined
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)
	fmt.Println("tx=", tx.Hash().Hex())
	elapsed := time.Since(start)
	log.Printf("\nContract deployment took %s\n", elapsed)

	// 7. Generate BLS parameters and a random G1 element
	y := func() {
		defer timeTrack(time.Now(), "y")
		pairing, err = pbc.NewPairingFromString(string(paramsByte))
		functions.CheckError(err)
		g = pairing.NewG1().Rand()
	}
	y()

	// 8. Call setParams(...) on the deployed contract
	z := func() {
		defer timeTrack(time.Now(), "z")
		tx1, err := instContract.SetParams(auth, paramsByte, g.Bytes())
		functions.CheckError(err)
		fmt.Println("\ntx1=", tx1.Hash().Hex())

		_, err = bind.WaitMined(ctx, connection, tx1)
		functions.CheckError(err)
	}
	z()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\nFunction %s took %s\n", name, elapsed)
}

func main() {
	// Ensure the tmp folder exists
	functions.Createtmp()

	// Generate BLS pairing parameters
	params := pbc.GenerateA(160, 512)

	// Deploy the contract and set its parameters
	CommContract([]byte(params.String()))
}
