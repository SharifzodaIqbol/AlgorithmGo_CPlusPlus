package main

import "fmt"

func kAnagram(s1 string, s2 string, k int) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr := [26]rune{}
	count := 0
	for _, char := range s1 {
		arr[char-'a']++
	}
	for _, char := range s2 {
		if arr[char-'a'] > 0 {
			arr[char-'a']--
		} else {
			count++
		}
		if count > k {
			return false
		}
	}
	return true
}

func main() {
	s1 := "geeks" // 1 1 0 -1 -1
	s2 := "eggkf" // 1 0 -1 -1
	k := 1
	fmt.Println(kAnagram(s1, s2, k))
}
