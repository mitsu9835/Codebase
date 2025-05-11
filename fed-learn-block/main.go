package main

import (
    "fmt"
    bls "github.com/kilic/bls12-381"
)

func main() {
    g1 := bls.NewG1()
    g2 := bls.NewG2()
    e := bls.NewEngine()

    p1 := g1.One()
    p2 := g2.One()
    e.AddPair(p1, p2)
    result := e.Result()

    fmt.Println("Pairing result:", result)
}
