package main

import "fmt"

func smallesWindow012(s string) int {
	i := 0
	for i < len(s)-2 {
		if s[i] == '0' && s[i+1] == '1' && s[i+2] == '2' {
			return 3
		}
		i++
	}
	return -1
}

func main() {
	s := "012"
	fmt.Println(smallesWindow012(s))
}
