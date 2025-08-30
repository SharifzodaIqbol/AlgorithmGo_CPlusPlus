package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func sentencePalindrome(scanner *bufio.Scanner) bool {
	scanner.Scan()
	text := scanner.Text()
	left, right := 0, len(text)-1

	for left < right {
		leftChar := unicode.ToLower(rune(text[left]))
		rightChar := unicode.ToLower(rune(text[right]))

		if !unicode.IsLetter(leftChar) && !unicode.IsDigit(leftChar) {
			left++
			continue
		}
		if !unicode.IsLetter(rightChar) && !unicode.IsDigit(rightChar) {
			right--
			continue
		}
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(sentencePalindrome(scanner))
}
