package main

import "fmt"

func removeDublicates(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			result = append(result, nums[i])
		}
	}
	return result
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 6}
	fmt.Println(removeDublicates(nums))
}
