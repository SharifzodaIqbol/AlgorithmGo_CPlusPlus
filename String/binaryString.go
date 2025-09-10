package main

import "fmt"

func isBinary(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i]-'0' != 1 && s[i]-'0' != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "01010101010b"
	fmt.Println(isBinary(s))
}
