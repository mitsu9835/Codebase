// functions/crypto.go
package functions

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "log"
    "math/big"

    bls12381 "github.com/kilic/bls12-381"
)

//
// ECDSA signature type
//
type ECDSASignature struct {
    R *big.Int
    S *big.Int
}

//
// MaskGradient adds element-wise noise to a gradient vector.
//
func MaskGradient(grad []float64, noise []float64) []float64 {
    n := len(grad)
    out := make([]float64, n)
    for i := 0; i < n; i++ {
        out[i] = grad[i] + noise[i]
    }
    return out
}

//
// CommitVector produces a BLS12-381 G1 commitment to a real vector:
//     COMIT(v) = Σ_i v[i] · G
//
func CommitVector(vec []float64) *bls12381.PointG1 {
    g1 := bls12381.NewG1()
    gen := g1.One()
    acc := g1.Zero()
    for _, v := range vec {
        bi := big.NewInt(int64(v))
        tmp := g1.New()
        g1.MulScalarBig(tmp, gen, bi)
        g1.Add(acc, acc, tmp)
    }
    return acc
}

//
// VerifyCommitments checks that
//    CommitVector(masked) == commitG + commitN
//
func VerifyCommitments(masked []float64, commitG, commitN *bls12381.PointG1) bool {
    commitM2 := CommitVector(masked)
    g1 := bls12381.NewG1()
    sum := g1.New()
    g1.Add(sum, commitG, commitN)
    return g1.Equal(sum, commitM2)
}

//
// SignCommitmentECDSA signs the SHA-256 hash of msg under priv,
// returning an ECDSASignature {R,S}.
//
func SignCommitmentECDSA(priv *ecdsa.PrivateKey, msg []byte) (*ECDSASignature, error) {
    hash := sha256.Sum256(msg)
    r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
    if err != nil {
        return nil, err
    }
    return &ECDSASignature{R: r, S: s}, nil
}

//
// VerifySignatureECDSA verifies an ECDSA signature (r,s) over SHA-256(msg)
// using pub.
//
func VerifySignatureECDSA(pub *ecdsa.PublicKey, msg []byte, r, s *big.Int) bool {
    hash := sha256.Sum256(msg)
    return ecdsa.Verify(pub, hash[:], r, s)
}

//
// GenerateKeyPair creates a new ECDSA P-256 private/public key pair.
//
func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Fatalf("GenerateKeyPair: %v", err)
    }
    return priv, &priv.PublicKey
}
