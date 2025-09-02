package main

import (
	"fmt"
)

func maxSumSub(arr []int, k int) int {
	if k >= len(arr) {
		fmt.Println("Invalid")
		return -1
	}
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum := windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		maxSum = max(maxSum, windowSum)
	}
	return maxSum
}

func main() {
	arr := []int{100, 200, 300, 400}
	k := 2
	fmt.Println(maxSumSub(arr, k))
}
