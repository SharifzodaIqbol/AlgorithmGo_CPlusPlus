package main

import (
	"fmt"
	"strings"
)

func findSubString(txt string, pat string) int {
	return strings.Index(txt, pat)
}

func main() {
	txt := "geeksforgeeks"
	pat := "eks"
	fmt.Println(findSubString(txt, pat))
}
