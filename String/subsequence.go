package main

import "fmt"

func isSubseq(s1 string, s2 string) bool {
	sizeS1 := len(s1)
	sizeS2 := len(s2)
	if sizeS1 > sizeS2 {
		return false
	}
	j, k := 0, 0
	for j < sizeS1 && k < sizeS2 {
		if s1[j] == s2[k] {
			j++
		}
		k++
	}
	return j == sizeS1
}

func main() {
	s1 := "AXY"
	s2 := "ADXCP"
	fmt.Println(isSubseq(s1, s2))
}
