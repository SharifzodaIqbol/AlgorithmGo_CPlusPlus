package main

import (
	"fmt"
)

func getSecondLargest(arr []int) int {
	firstLarge, secondLarge := -1, -1
	for _, value := range arr {
		if value > firstLarge {
			firstLarge = value
		}
	}
	for _, value := range arr {
		if value > secondLarge && value != firstLarge {
			secondLarge = value
		}
	}
	return secondLarge
}
func main() {
	arr := []int{10, 10, 10, 1, 1}
	fmt.Println(getSecondLargest(arr))
}
