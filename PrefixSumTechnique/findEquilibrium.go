package main

import "fmt"

func findEquilibrium(arr []int) int {
	size := len(arr)
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}
	leftSum := 0
	for i := 0; i < size; i++ {
		rightSum := totalSum - leftSum - arr[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += arr[i]
	}
	return -1
}
func main() {
	arr := []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Println(findEquilibrium(arr))
}
