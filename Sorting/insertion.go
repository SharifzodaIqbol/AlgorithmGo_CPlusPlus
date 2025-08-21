package main

import "fmt"

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}
func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println(insertionSort(arr))
}
