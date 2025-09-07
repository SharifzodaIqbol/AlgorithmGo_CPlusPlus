package main

import (
	"fmt"
)

func splitArray(arr []int) ([]int, []int) {
	totalSum := 0

	for _, num := range arr {
		totalSum += num
	}
	currentSum := 0
	indx := -1
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		totalSum -= arr[i]
		if totalSum == currentSum {
			indx = i + 1
			break
		}
	}
	firstPart := make([]int, indx)
	secondPart := make([]int, len(arr)-indx)
	if indx == -1 {
		fmt.Println("Not Possible")
		return firstPart, secondPart
	}
	copy(firstPart, arr[:indx])
	copy(secondPart, arr[indx:])
	return firstPart, secondPart
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(splitArray(arr))
}
