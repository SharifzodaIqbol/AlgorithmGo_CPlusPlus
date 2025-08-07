package main

import (
	"fmt"
)

func Prime(n int) bool {
	result := 2
	for result*result <= n {
		if n%result == 0 {
			return false
		}
		result++
	}
	return true
}
func main() {
	var n int
	fmt.Scan(&n)
	for i := 2; i*i <= n; i++ {
		if Prime(i) {
			fmt.Print(i*i, " ")
		}
	}
}
