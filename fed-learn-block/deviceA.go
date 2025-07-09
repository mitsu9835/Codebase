package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/csv"
	"encoding/hex"
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
	const workspace = "tmp/fog1"
	deviceFile := fmt.Sprintf("%s/device_output.json", workspace)
	pubFile := fmt.Sprintf("%s/fog1_pubkey.json", workspace)
	sigFile := fmt.Sprintf("%s/fog1_output.json", workspace)

	// 1) Load CSV
	f, err := os.Open("train.csv")
	if err != nil {
		log.Fatalf("open train.csv: %v", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	// 2) Header → feature count
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("read header: %v", err)
	}
	n := len(header) - 1
	fmt.Printf("Device A: Detected %d features (excluding target)\n", n)

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
	fmt.Printf("Device A: Total examples: %d\n", m)

	// 4) Parse X, y
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

	// 5) Init θ = [0.1, 0.2, ...]
	theta := make([]float64, n)
	for i := range theta {
		theta[i] = 0.1 * float64(i+1)
	}
	fmt.Printf("Device A: Initialized theta: %v\n", theta)

	// 6) Dot prod on first example
	dot := 0.0
	for j := 0; j < n; j++ {
		dot += theta[j] * X[0][j]
	}
	fmt.Printf("Device A: Dot product for example #1: %f\n", dot)

	// 7) Compute local gradient
	grad := computeLocalGradient(X, y, theta)
	fmt.Printf("Device A: Local gradient: %v\n", grad)

	// ── On-chain noise ──

	// 8) Load contract address
	addrBytes, err := os.ReadFile("tmp/contractInfo/contract-address")
	if err != nil {
		log.Fatalf("read address: %v", err)
	}
	addr := common.HexToAddress(strings.TrimSpace(string(addrBytes)))
	fmt.Printf("Device A: Using Store at %s\n", addr.Hex())

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("dial node: %v", err)
	}
	store, err := contracts.NewContracts(addr, client)
	if err != nil {
		log.Fatalf("bind contract: %v", err)
	}

	// 9) Fetch noise per feature
	noises := make([]float64, n)
	for j := 0; j < n; j++ {
		nb, err := store.GenerateNoise(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			log.Fatalf("generate_noise[%d]: %v", j, err)
		}
		noises[j] = float64(nb.Int64())
	}
	fmt.Printf("Device A: Noise vector η: %v\n", noises)

	// 10) Mask gradient
	masked := functions.MaskGradient(grad, noises)
	fmt.Printf("Device A: Masked gradient: %v\n", masked)

	// 11) Commitments
	g1 := bls.NewG1()
	commitG := functions.CommitVector(grad)
	commitN := functions.CommitVector(noises)
	commitM := functions.CommitVector(masked)
	cgBytes := g1.ToCompressed(commitG)

	// 12) Prepare the workspace & clear any old fog files
	if err := os.MkdirAll(workspace, 0755); err != nil {
		log.Fatalf("mkdir %s: %v", workspace, err)
	}
	_ = os.Remove(deviceFile)
	_ = os.Remove(pubFile)
	_ = os.Remove(sigFile)

	// 13) Write exactly device_output.json
	out := struct {
		MaskedUpdate []float64 `json:"masked_update"`
		CommitG      string    `json:"commit_g"`
		CommitN      string    `json:"commit_n"`
		CommitM      string    `json:"commit_m"`
	}{
		MaskedUpdate: masked,
		CommitG:      fmt.Sprintf("%x", cgBytes),
		CommitN:      fmt.Sprintf("%x", g1.ToCompressed(commitN)),
		CommitM:      fmt.Sprintf("%x", g1.ToCompressed(commitM)),
	}
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		log.Fatalf("marshal JSON: %v", err)
	}
	if err := ioutil.WriteFile(deviceFile, data, 0644); err != nil {
		log.Fatalf("write %s: %v", deviceFile, err)
	}
	fmt.Printf("Device A: Wrote %s\n", deviceFile)

	// 14) Wait for fog1 to emit pubkey+signature
	fmt.Printf("Device A: Waiting for %s …\n", sigFile)
	for {
		if _, err := os.Stat(sigFile); err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// 15) Load fresh pubkey
	pubData, err := ioutil.ReadFile(pubFile)
	if err != nil {
		log.Fatalf("read %s: %v", pubFile, err)
	}
	var pubJSON struct{ X, Y string }
	if err := json.Unmarshal(pubData, &pubJSON); err != nil {
		log.Fatalf("unmarshal %s: %v", pubFile, err)
	}
	xB, _ := hex.DecodeString(pubJSON.X)
	yB, _ := hex.DecodeString(pubJSON.Y)
	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(xB),
		Y:     new(big.Int).SetBytes(yB),
	}

	// 16) Load fresh signature
	sigData, err := ioutil.ReadFile(sigFile)
	if err != nil {
		log.Fatalf("read %s: %v", sigFile, err)
	}
	var sigJSON struct{ R, S string }
	if err := json.Unmarshal(sigData, &sigJSON); err != nil {
		log.Fatalf("unmarshal %s: %v", sigFile, err)
	}
	rB, _ := hex.DecodeString(sigJSON.R)
	sB, _ := hex.DecodeString(sigJSON.S)
	r := new(big.Int).SetBytes(rB)
	s := new(big.Int).SetBytes(sB)

	// 17) Verify Fog’s signature over this round’s commitG
	if functions.VerifySignatureECDSA(pubKey, cgBytes, r, s) {
		fmt.Println("Device A: Fog signature valid ✅")
	} else {
		fmt.Println("Device A: Fog signature invalid ❌")
	}
}
