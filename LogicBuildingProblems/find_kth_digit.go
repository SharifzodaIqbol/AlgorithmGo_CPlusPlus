package main

import (
	"fmt"
	"math/big"
)

func aDegreeb(a, b int) *big.Int {
	result := big.NewInt(1)
	bigA := big.NewInt(int64(a))
	for i := 0; i < b; i++ {
		result.Mul(result, bigA)
	}
	return result
}
func find_kth_digit(a, b, k int) string {
	pow := aDegreeb(a, b).String()
	if k > len(pow) {
		return "k is too large"
	}
	return string(pow[len(pow)-k])
}

func main() {
	a, b, k := 5, 2, 2
	fmt.Println(find_kth_digit(a, b, k))
}
