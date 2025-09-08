package main

import (
	"fmt"
)

func insertChar(s string, x rune, pos int) string {
	return s[:pos] + string(x) + s[pos:]
}

func main() {
	str := "HelloWorld"
	fmt.Println(insertChar(str, ' ', 5))
}
