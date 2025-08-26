package main

import "math"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func findClosest(arr1, arr2 []int, x int) []int {
	left := 0
	right := len(arr2) - 1
	mn := math.MinInt32
	for left <= right {
		sum := arr1[left] + arr2[right]
		if sum < x {
			left++
		}
	}
}

func main() {

}
