package main

import "fmt"

func orginalArray(presum []int) []int {
	for i := len(presum) - 1; i > 0; i-- {
		presum[i] = presum[i] - presum[i-1]
	}
	return presum
}

func main() {
	arr := []int{45, 57, 63, 78, 89, 97}
	fmt.Println(orginalArray(arr))
}
