package main

import "fmt"

func sort012(arr []int) int {
	indexL := 0
	//indexMid := indexR / 2
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[indexL] = arr[i]
			indexL++
		} else if arr[i] == 1 {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return indexL
}

func main() {
	arr := []int{0, 1, 2, 0, 1, 2}
	sort012(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
