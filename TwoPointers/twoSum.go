package main

import (
	"fmt"
	"slices"
)

func twoSum(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1
	slices.Sort(arr)
	for left < right {
		if arr[left]+arr[right] > target {
			right--
		} else if arr[left]+arr[right] < target {
			left++
		} else {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{0, -1, 2, -3, 1} // -3 -1 0 1 2
	target := -2
	fmt.Println(twoSum(arr, target))
}
