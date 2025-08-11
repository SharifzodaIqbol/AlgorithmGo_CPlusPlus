package main

import "fmt"

func reverse(nums []int) []int {
	size := len(nums)
	for i := 0; i < size/2; i++ {
		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
	}
	return nums
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(reverse(nums))
}
