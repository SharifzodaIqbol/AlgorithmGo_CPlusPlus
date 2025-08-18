package main

import "fmt"

func twoSum(arr []int, target int) bool {
	myMap := make(map[int]struct{})
	for _, value := range arr {
		if _, ok := myMap[value-target]; ok {
			return true
		}
		myMap[value] = struct{}{}
	}
	return false
}
func main() {
	arr := []int{0, -1, 2, -3, 1}
	target := -2
	fmt.Println(twoSum(arr, target))
}
