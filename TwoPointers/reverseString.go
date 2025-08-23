package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	runes := []rune(text)
	right := len(runes) - 1
	left := 0
	for left < right {
		if runes[left] == ' ' {
			left++
			continue
		} else if runes[right] == ' ' {
			right--
			continue
		} else {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}
	return string(runes)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(reverseString(scanner))
}
