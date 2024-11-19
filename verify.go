package main

import "math/big"


func verify(message int, signature *big.Int) bool {

	M := big.NewInt(int64(message))

	E, _ := new(big.Int).SetString(e, 10)
	
	N, _ := new(big.Int).SetString(n, 10)
	
	result := new(big.Int).Exp(signature, E, N)
	
	if result.Cmp(M) == 0 {
		return true
	}else {
		return false
	}
}