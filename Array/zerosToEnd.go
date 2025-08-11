package main

import "fmt"

func zerosToEnd(nums []int) []int {
	size := len(nums)
	count := 0
	for i := 0; i < size; i++ {
		if nums[i] != 0 {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	return nums
}

func main() {
	nums := []int{1, 2, 0, 4, 3, 0, 5, 0}
	fmt.Println(zerosToEnd(nums))
}
