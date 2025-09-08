package main

import (
	"fmt"
	"strings"
)

func searchX(s string, x rune) int {
	return strings.IndexRune(s, x)
}
func main() {
	s := "hello"
	x := 'l'
	fmt.Println(searchX(s, x))
}
