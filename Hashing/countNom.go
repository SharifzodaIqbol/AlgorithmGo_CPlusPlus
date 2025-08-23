package main

import (
	"fmt"
	"math"
)

func countNum(arr []int) int {
	Hash := make(map[int]struct{})
	mn := math.MaxInt32
	mx := math.MinInt32
	for _, value := range arr {
		if value < mn {
			mn = value
		}
		if value > mx {
			mx = value
		}
		Hash[value] = struct{}{}
	}
	return (mx - mn + 1) - len(Hash)
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(countNum(arr))
}
