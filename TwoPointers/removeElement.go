package main

import "fmt"

func removeElement(arr []int, ele int) int {
	indx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != ele {
			arr[indx] = arr[i]
			indx++
		}
	}
	return indx
}

func main() {
	arr := []int{0, 1, 3, 0, 2, 2, 4, 2}
	ele := 2
	fmt.Println(removeElement(arr, ele))
}
