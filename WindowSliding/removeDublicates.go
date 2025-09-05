package main

import "fmt"

func removeDublicates(s string) string {
	if len(s) == 0 {
		return ""
	}
	str := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			str += string(s[i+1])
		}
	}
	return str
}

func main() {
	s := "aabccba"
	fmt.Println(removeDublicates(s))
}
