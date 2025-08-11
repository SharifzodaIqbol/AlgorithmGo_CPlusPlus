package main

import "fmt"

func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
func main() {
	nums := []int{90, 80, 100, 70, 40, 30}
	fmt.Println(isSorted(nums))
}
