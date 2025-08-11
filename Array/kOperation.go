package main

import (
	"fmt"
	"math"
)

func maxArray(nums []int) int {
	mx := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] >= mx {
			mx = nums[i]
		}
	}
	return mx
}
func check(nums []int, k int) int {
	mx := maxArray(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if (mx-nums[i])%k == 0 {
			count += (mx - nums[i]) / k
		} else {
			return -1
		}
	}
	return count
}
func main() {
	nums := []int{4, 4, 4, 4}
	k := 3
	fmt.Println(check(nums, k))
	return
}
