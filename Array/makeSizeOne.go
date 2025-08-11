package main

import (
	"fmt"
	"math"
)

func minArray(nums []int) int {
	mn := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= mn {
			mn = nums[i]
		}
	}
	return mn
}
func minimumCost(nums []int) int {
	mn := minArray(nums)
	return (len(nums) - 1) * mn
}
func main() {
	nums := []int{3, 4}
	fmt.Println(minimumCost(nums))
}
