package main

import "fmt"

func findMissing(arr []int, low, high int) []int {
	slice := []int{}
	Hash := make(map[int]struct{})

	for _, value := range arr {
		Hash[value] = struct{}{}
	}
	for i := low; i <= high; i++ {
		if _, exist := Hash[i]; !exist {
			slice = append(slice, i)
		}
	}
	return slice
}

func main() {
	arr := []int{1, 14, 11, 51, 15}
	low, high := 50, 55
	fmt.Println(findMissing(arr, low, high))
}
