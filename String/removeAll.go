package main

import (
	"fmt"
)

func removeAllChar(s []rune, ch rune) string {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ch {
			s[j] = s[i]
			j++
		}
	}
	return string(s[:j])
}

func main() {
	s := []rune("HelloWorld")
	fmt.Println(removeAllChar(s, 'l'))
}
