package main

import "fmt"

func minRemove(arr1, arr2 []int) int {
	Hash := make(map[int]int)
	for _, value := range arr1 {
		Hash[value]++
	}
	count := 0
	for _, value := range arr2 {
		if _, exist := Hash[value]; exist {
			if Hash[value] > 0 {
				Hash[value]--
			}
			count++
		}
	}
	return count
}

func main() {
	arr1 := []int{4, 2, 4, 4}
	arr2 := []int{4, 3}
	fmt.Println(minRemove(arr1, arr2))
}
