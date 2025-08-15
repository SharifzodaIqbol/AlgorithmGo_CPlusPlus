package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	slice := []int{3, 1, 3}
	N := len(slice)
	sumN := N * (N + 1) / 2
	sumSlice := 0
	dubl := 0
	for _, value := range slice {
		if _, ok := myMap[value]; ok {
			dubl = value
		}
		sumSlice += value
		myMap[value]++
	}
	miss := dubl - (sumSlice - sumN)
	fmt.Println(dubl, miss)
}
