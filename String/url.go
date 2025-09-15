package main

import (
	"bufio"
	"fmt"
	"strings"
)

func urlIfy(s *bufio.Scanner) string {
	s.Scan()
	text := s.Text()
	return strings.ReplaceAll(text, " ", "%20")
}

func main() {
	s := bufio.NewScanner(strings.NewReader("i love programming"))
	fmt.Println(urlIfy(s))
}
