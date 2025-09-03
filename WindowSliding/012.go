package main

import (
	"fmt"
	"math"
)

func window012(s string) int {
	var arr [3]int
	var left int
	var minLenght int = math.MaxInt32
	for right := 0; right < len(s); right++ {
		char := s[right] - '0'
		arr[char]++
		for arr[0] > 0 && arr[1] > 0 && arr[2] > 0 {
			windowSize := right - left
			if windowSize < minLenght {
				minLenght = windowSize
			}
			arr[s[left]-'0']--
			left++
		}
	}
	if minLenght == math.MaxInt32 {
		return -1
	}
	return minLenght + 1
}

func main() {
	s := "012"
	fmt.Println(window012(s))
}
