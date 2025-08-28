package main

import (
	"fmt"
	"slices"
)

func dominatePair(arr []int) int {
	size := len(arr)
	count := 0
	slices.Sort(arr[:size/2])
	slices.Sort(arr[size/2:])
	right := size / 2
	for left := 0; left < size/2; left++ {
		for right < size && arr[left] >= 5*arr[right] {
			right++
		}
		count += (right - size/2)
	}
	return count
}

func main() {
	arr := []int{10, 2, 2, 1} // 2 10 1 2
	fmt.Println(dominatePair(arr))
}
