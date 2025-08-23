package main

import "fmt"

// from 1 to n-1
func onlyRepeating(arr []int) int {
	Hash := make(map[int]int)

	for _, value := range arr {
		if _, exist := Hash[value]; exist {
			return value
		}
		Hash[value]++
	}
	return -1
}

func main() {
	arr := []int{1, 3, 2, 3, 4}
	fmt.Println(onlyRepeating(arr)) // 3
}
