package main

import "fmt"

func pairSumInRotatedArray(arr []int, target int) bool {
	n := len(arr)
	i := 0
	for ; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			break
		}
	}
	left := (i + 1) % n
	right := i
	for left != right {
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		} else if sum < target {
			left = (left + 1) % n
		} else {
			right = (right - 1 + n) % n
		}
	}
	return false
}

func main() {
	arr := []int{11, 15, 6, 8, 9, 10}
	target := 16
	fmt.Println(pairSumInRotatedArray(arr, target))
}
