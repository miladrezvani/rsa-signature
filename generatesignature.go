package main

import "math/big"


func generatesignature(signature1, signature2 *big.Int) *big.Int {

	signature3 := new(big.Int)

    signature3.Mul(signature1, signature2)

	one, _ := new(big.Int).SetString("1", 10)

	N, _ := new(big.Int).SetString(n, 10)

	result := new(big.Int).Exp(signature3, one, N)

    return result
}