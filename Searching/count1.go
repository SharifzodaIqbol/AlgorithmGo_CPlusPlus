package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 0}
	left, right := 0, len(arr)-1
	target := 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left)
}
