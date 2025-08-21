package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
func main() {
	arr := []int{12, 11, 13, 5, 6} // 5 6 11 12 13
	fmt.Println(bubbleSort(arr))
}
