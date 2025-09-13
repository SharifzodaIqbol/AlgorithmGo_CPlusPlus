package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func check(arr [26]rune) bool {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == 0 {
			return false
		}
	}
	return true
}

func isPangram(s *bufio.Scanner) bool {
	arr := [26]rune{}
	s.Scan()
	text := s.Text()
	for i := 0; i < len(text); i++ {
		toRune := unicode.ToLower(rune(text[i]))
		if unicode.IsLetter(toRune) {
			arr[toRune]++
		}
	}
	if check(arr) {
		return true
	}
	return false
}

func main() {
	s := bufio.NewScanner(strings.NewReader("The quick brown fox jumps over the lazy dog"))
	fmt.Println(isPangram(s))
}
