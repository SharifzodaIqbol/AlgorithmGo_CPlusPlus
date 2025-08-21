package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println(selectionSort(arr)) //11 12 22 25 64
}
