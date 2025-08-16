package main

import "fmt"

//Only Repeating From 1 To n-1
func onlyRepeating(arr []int) int {
	N := len(arr) - 1
	sumN := N * (N + 1) / 2
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum - sumN
}
func main() {
	arr := []int{1, 5, 1, 2, 3, 4}
	fmt.Println(onlyRepeating(arr))
}
