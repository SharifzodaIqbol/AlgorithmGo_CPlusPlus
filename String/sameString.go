package main

import "fmt"

func isEqual(str1, str2 string) bool {
	return str1 == str2
}

func main() {
	str1 := "str1"
	str2 := "str2"
	fmt.Println(isEqual(str1, str2))
}
