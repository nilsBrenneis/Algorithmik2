package main

import (
	"math"
)

func isPrime(value int) (bool, int) {
	n := float64(value)
	prim := true
	k := 2.0
	for k < n {
		if math.Mod(n, k) == 0 {
			prim = false
		}
		k++
	}
	return prim, value
}
