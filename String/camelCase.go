package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func camelCase(scanner *bufio.Scanner) string {
	scanner.Scan()
	s := scanner.Text()
	runes := make([]rune, len(s))
	j := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			flag = true
			continue
		} else if flag {
			runes[j] = unicode.ToUpper(rune(s[i]))
			flag = false
			j++
		} else {
			runes[j] = rune(s[i])
			j++
		}
	}
	return string(runes)
}

func main() {
	str := bufio.NewScanner(strings.NewReader("here comes the garden"))
	fmt.Println(camelCase(str))
}
