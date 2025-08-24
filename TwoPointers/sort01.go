package main

import "fmt"

func sort01(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		for left < right && arr[left] != 1 {
			left++
		}
		for left < right && arr[right] != 0 {
			right--
		}
		if left >= right {
			break
		}
		arr[right], arr[left] = arr[left], arr[right]
		left++
		right--
	}
	return arr
}

func main() {
	arr := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0}
	fmt.Println(sort01(arr))
}
