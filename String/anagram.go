package main

import (
	"fmt"
)

func isAnagram(s1 string, s2 string) bool {
	size1 := len(s1)
	size2 := len(s2)
	if size1 != size2 {
		return false
	}
	arr := [26]rune{}

	for _, char := range s1 {
		arr[char-'a']++
	}
	for _, char := range s2 {
		arr[char-'a']--
	}
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "geeks"
	s2 := "kseeg"
	fmt.Println(isAnagram(s1, s2))
}
