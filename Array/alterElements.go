package main

import "fmt"

func alterElem(nums []int, count int, result []int) []int {
	if count >= len(nums) {
		return result
	}
	result = append(result, nums[count])
	return alterElem(nums, count+2, result)
}
func main() {
	arr := []int{10, 20, 30, 40, 50}
	result := []int{}
	fmt.Println(alterElem(arr, 0, result))
}
