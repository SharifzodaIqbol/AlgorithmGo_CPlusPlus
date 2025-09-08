package main

import "fmt"

func removeChar(s string, pos int) string {
	return s[:pos] + s[pos+1:]
}
func main() {
	str := "abcde"
	fmt.Println(removeChar(str, 2))
}
