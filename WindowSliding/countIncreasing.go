package main

import "fmt"

func countIncreasing(arr []int) int {
	size := len(arr)
	count := 0
	result := 0
	for i := 0; i < size-1; i++ {
		if arr[i] < arr[i+1] {
			count++
		} else {
			result += count * (count + 1) / 2
			count = 0
		}
	}
	result += count * (count + 1) / 2
	return result
}

func main() {
	arr := []int{1, 4, 5, 3, 7, 9}
	fmt.Println(countIncreasing(arr))
}

/// 1 2 3 4 // 1 2 3 4| 1 2 3| 1 2| 2 3 4| 2 3| 3 4|
