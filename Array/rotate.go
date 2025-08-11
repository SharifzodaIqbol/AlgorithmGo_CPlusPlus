package main

import "fmt"

func reverse[T any](nums []T, begin, end int) []T {
	for i := begin; i < end/2; i++ {
		nums[i], nums[end-i-1] = nums[end-i-1], nums[i]
	}
	return nums
}
func rotate[T any](nums []T, d int) []T {
	size := len(nums)
	d %= size
	reverse(nums, 0, size)
	reverse(nums, 0, d)
	reverse(nums, d, size+d)
	return nums
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	d := 7
	fmt.Println(rotate(nums, d)) // 4 5 1 2 3

}
