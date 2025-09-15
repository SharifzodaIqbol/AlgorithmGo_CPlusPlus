package main

import "fmt"

func isPolindrom(s string) bool {
	size := len(s)
	for i := 0; i < size/2; i++ {
		if s[i] != s[size-i-1] {
			return false
		}
	}
	return true
}

/*
func isPolindrom(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
*/
func main() {
	s := "abba"
	fmt.Println(isPolindrom(s))
}
