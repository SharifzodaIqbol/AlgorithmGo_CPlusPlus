package main

import (
	"fmt"
	"math"
)

func minSubsets(arr []int) int {
	Hash := make(map[int]int)
	mx := math.MinInt32
	for _, value := range arr {
		Hash[value]++
		if Hash[value] > mx {
			mx = Hash[value]
		}
	}
	return mx
}

func main() {
	arr := []int{40, 50, 30, 40, 50, 30, 30}
	fmt.Println(minSubsets(arr))
}
