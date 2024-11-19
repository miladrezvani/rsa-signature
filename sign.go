package main

import "math/big"


func sign(message int) *big.Int {

	M := big.NewInt(int64(message))
   
	D, _ := new(big.Int).SetString(d, 10)
   
	N, _ := new(big.Int).SetString(n, 10)
   
	result := new(big.Int).Exp(M, D, N)
   
	return result
}