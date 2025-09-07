package main

import "fmt"

func findMean(arr []int, queries [][]int) []int {
	result := []int{}
	arrSum := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		arrSum[i] = arrSum[i-1] + arr[i-1]
	}
	for i := 0; i < len(queries); i++ {
		l := queries[i][0] - 1
		r := queries[i][1] - 1
		sum := arrSum[r+1] - arrSum[l]
		count := r - l + 1
		result = append(result, sum/count)
	}
	return result
}

func main() {
	arr := []int{3, 7, 2, 8, 5}
	queries := [][]int{{1, 3}, {2, 5}}
	fmt.Println(findMean(arr, queries))
}
