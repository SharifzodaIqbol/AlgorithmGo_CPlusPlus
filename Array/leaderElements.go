package main

import (
	"fmt"
	"math"
)

func leaderElements(nums []int) {
	mx := math.MinInt32

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= mx {
			mx = nums[i]
			defer fmt.Printf("%d ", nums[i])
		}
	}
}
func main() {
	nums := []int{16, 17, 4, 3, 5, 2}
	leaderElements(nums)
}
