package main

import (
	"fmt"
	"math"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func findClosest(arr1, arr2 []int, x int) []int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return []int{}
	}
	left := 0
	right := len(arr2) - 1
	leftElem, rightElem := 0, 0
	mn := math.MaxInt32
	for left <= len(arr1)-1 && right >= 0 {
		sum := arr1[left] + arr2[right]
		if abs(sum-x) < mn {
			mn = abs(sum - x)
			leftElem, rightElem = arr1[left], arr2[right]
		} else if sum < x && left < len(arr1)-1 {
			left++
		} else if sum > x || left == len(arr1)-1 {
			right--
		}
		if abs(sum-x) > mn {
			break
		}
	}
	return []int{leftElem, rightElem}
}

// arr2[] = {1, 2, 3, 4};
// arr1[] = {4, 3, 2, 1};
// traget := 8
func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 3, 2, 1}
	x := 50
	fmt.Println(findClosest(arr1, arr2, x))
}
