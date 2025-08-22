package main

import (
	"fmt"
	"math"
)

func mostFreqElem(arr []int) int {
	Hash := make(map[int]int)
	mx := 0
	maxElem := math.MinInt32
	for _, value := range arr {
		Hash[value]++
		mx = max(mx, Hash[value])
	}
	for key, value := range Hash {
		if value == mx {
			maxElem = max(maxElem, key)
		}
	}
	return maxElem
}

func main() {
	arr := []int{1, 2, 2, 4, 1}
	fmt.Println(mostFreqElem(arr))
}
