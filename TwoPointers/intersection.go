package main

import (
	"fmt"
	"slices"
)

func binarySearch(arr1 []int, target int) bool {
	left := 0
	right := len(arr1) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr1[mid] == target {
			return true
		} else if arr1[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
func intersectionTwoArray(arr1, arr2 []int) {
	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	slices.Sort(arr1)
	for i := 0; i < len(arr2); i++ {
		if binarySearch(arr1, arr2[i]) {
			fmt.Printf("%d ", arr2[i])
		}
	}
}
func main() {
	arr1 := []int{5, 6, 2, 1, 4}
	arr2 := []int{7, 9, 4, 2}
	intersectionTwoArray(arr1, arr2)
}
