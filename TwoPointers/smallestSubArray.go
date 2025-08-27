package main

import (
	"fmt"
	"math"
)

func smallestSubArray(x int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	left := 0
	mn := math.MaxInt32
	sum := 0
	for next := 0; next < len(arr); next++ {
		sum += arr[next]
		for sum > x && left <= next {
			if length := next - left + 1; length < mn {
				mn = length
			}
			sum -= arr[left]
			left++
		}
	}
	if mn == math.MaxInt32 {
		return 0
	}
	return mn
}

func main() {
	x := 100
	arr := []int{1, 10, 5, 2, 7}
	fmt.Println(smallestSubArray(x, arr))
}
