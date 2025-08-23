package main

import "fmt"

func twoSumaPair(arr []int, target int) int {
	Hash := make(map[int]int)
	count := 0
	for _, value := range arr {
		if valueHash, ok := Hash[target-value]; ok {
			count += valueHash
		}
		Hash[value]++
	}
	return count
}

func main() {
	arr := []int{1, 5, 7, -1, 5}
	target := 6
	fmt.Println(twoSumaPair(arr, target))
}
