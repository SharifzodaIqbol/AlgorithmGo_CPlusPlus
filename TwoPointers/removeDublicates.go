package main

import "fmt"

func removeDeblicates(arr []int) int {
	size := len(arr)
	if size <= 1 {
		return size
	}
	index := 1
	for i := 1; i < size; i++ {
		if arr[i-1] != arr[i] {
			arr[index] = arr[i]
			index++
		}
	}
	return index
}

func main() {
	arr := []int{1, 2, 3}
	newSize := removeDeblicates(arr)
	for i := 0; i < newSize; i++ {
		fmt.Println(arr[i])
	}
}
