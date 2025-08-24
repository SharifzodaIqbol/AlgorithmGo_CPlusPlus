package main

import "fmt"

func sort012(arr []int) []int {
	indexL := 0
	indexMid := 0
	indexR := len(arr) - 1
	for indexMid <= indexR {
		if arr[indexMid] == 0 {
			arr[indexL], arr[indexMid] = arr[indexMid], arr[indexL]
			indexL++
			indexMid++
		} else if arr[indexMid] == 1 {
			indexMid++
		} else {
			arr[indexMid], arr[indexR] = arr[indexR], arr[indexMid]
			indexR--
		}
	}
	return arr
}

func main() {
	arr := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	fmt.Println(sort012(arr))
}
