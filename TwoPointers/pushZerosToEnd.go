package main

import "fmt"

func pushZerosToEnd(arr []int) {
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[k], arr[i] = arr[i], arr[k]
			k++
		}
	}
}

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	pushZerosToEnd(arr)
	fmt.Println(arr)
}
