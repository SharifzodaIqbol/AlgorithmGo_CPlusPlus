package main

import (
	"fmt"
	"slices"
)

func mergeArrays(arr1 []int, arr2 []int) {
	n1, n2 := len(arr1), len(arr2)
	res := make([]int, 0, n1+n2)
	res = append(res, arr1...)
	res = append(res, arr2...)
	slices.Sort(res)
	copy(arr1, res[:n1])
	copy(arr2, res[n1:])
}

func main() {
	arr1 := []int{1, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8}
	mergeArrays(arr1, arr2)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
