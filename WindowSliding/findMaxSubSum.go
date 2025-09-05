package main

import "fmt"

func findMaxSubSum(arr []int, sum int) int {
	mx := 0
	result := arr[0]
	start := 0
	for i := 1; i < len(arr); i++ {
		if result <= sum {
			mx = max(mx, result)
		}
		for start < i && result+arr[i] > sum {
			result -= arr[start]
			start++
		}
		if result < 0 {
			result = 0
		}
		result += arr[i]
	}
	if result <= sum {
		mx = max(mx, result)
	}
	return mx
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 11
	fmt.Println(findMaxSubSum(arr, sum))
}
