package main

import "fmt"

func binarySubString(s string) int {
	count := -1
	sum := 0
	for _, char := range s {
		if char == '1' {
			count++
			sum += count
		}
	}
	return sum
}

func main() {
	str := "111"
	fmt.Println(binarySubString(str))
}
