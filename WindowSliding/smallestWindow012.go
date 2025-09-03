package main

import (
	"fmt"
	"math"
)

func smallesWindow012(s string) int {
	zero, one, two := false, false, false
	indexZero, indexOne, indexTwo := 0, 0, 0
	result := math.MaxInt32
	for i, value := range s {
		if value == '0' {
			zero = true
			indexZero = i
		} else if value == '1' {
			one = true
			indexOne = i
		} else if value == '2' {
			two = true
			indexTwo = i
		}

		if zero && one && two {
			result = min(result,
				max(indexZero, indexOne, indexTwo)-min(indexZero, indexOne, indexTwo))
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result + 1
}

func main() {
	s := "0012"
	fmt.Println(smallesWindow012(s))
}
