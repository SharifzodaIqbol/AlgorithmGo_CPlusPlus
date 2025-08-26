package main

import (
	"fmt"
	"math"
	"slices"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func sumClosest(arr []int, target int) []int {
	if len(arr) < 2 {
		return []int{}
	}
	slices.Sort(arr)
	left := 0
	right := len(arr) - 1
	mn := math.MaxInt32
	x, y := 0, 0
	for left < right {
		sum := arr[left] + arr[right]
		if abs(sum-target) < mn {
			mn = abs(sum - target)
			x, y = arr[left], arr[right]
		} else if sum == target {
			return []int{arr[left], arr[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{x, y}
}

func main() {
	arr := []int{10} // 1 2 4 5 7
	target := 10
	fmt.Println(sumClosest(arr, target))

}
