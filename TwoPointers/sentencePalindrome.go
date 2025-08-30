package main

import (
	"bufio"
	"fmt"
	"os"
)

func sentencePalindrome(scanner *bufio.Scanner) bool {
	scanner.Scan()
	text := scanner.Text()
	runes := []rune(text)
	left, right := 0, len(text)-1
	for left < right {
		if runes[left] >= 65 && runes[left] <= 90 {
			runes[left] += 32
		}
		if runes[right] >= 65 && runes[right] <= 90 {
			runes[right] += 32
		}
		if (runes[left] < 48 || runes[left] > 57) && (runes[left] < 97 || runes[left] > 122) {
			left++
			continue
		}
		if (runes[right] < 48 || runes[right] > 57) && (runes[right] < 97 || runes[right] > 122) {
			right--
			continue
		}
		if runes[left] != runes[right] {
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
