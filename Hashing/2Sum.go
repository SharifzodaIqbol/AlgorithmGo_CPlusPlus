package main

import "fmt"

func twoSumaPair(arr []int, target int) bool {
	Hash := make(map[int]struct{})
	for _, value := range arr {
		if _, ok := Hash[target-value]; ok {
			return true
		}
		Hash[value] = struct{}{}
	}
	return false
}

func main() {
	arr := []int{0, -1, 2, -3, 1}
	target := 20
	fmt.Println(twoSumaPair(arr, target))
}
