package main

import "fmt"

func reverse(s string) string {
	n := len(s)
	sCopy := []rune(s)
	for i := 0; i < n/2; i++ {
		sCopy[i], sCopy[n-i-1] = sCopy[n-i-1], sCopy[i]
	}
	return string(sCopy)
}

func main() {
	a := "ab"
	fmt.Println(reverse(a))
}
