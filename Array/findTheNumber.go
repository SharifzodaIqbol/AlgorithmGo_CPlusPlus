package main

import "fmt"

//Given an array arr[] of size n-1 with distinct integers in the range of [1, n].
func findTheMissingNumber(slice []int) int {
	n := len(slice) + 1
	sumSlice := 0
	for i := 0; i < len(slice); i++ {
		sumSlice += slice[i]
	}
	result := ((n * (n + 1)) / 2) - sumSlice
	return result
}
func main() {
	slice := []int{8, 2, 4, 5, 3, 7, 1}
	fmt.Println(findTheMissingNumber(slice))
}
