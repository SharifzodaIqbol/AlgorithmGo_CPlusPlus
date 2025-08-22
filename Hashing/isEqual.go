package main

import "fmt"

func isEqual(a []int, b []int) bool {
	Hash := make(map[int]int)

	for _, value := range a {
		Hash[value]++
	}
	for _, value := range b {
		if _, ok := Hash[value]; !ok {
			return false
		}
		if Hash[value] <= 0 {
			return false
		}
		Hash[value]--
	}
	for key, value := range Hash {
		fmt.Println(key, value)
	}
	return true
}

func main() {
	a := []int{3, 5, 2, 5, 2}
	b := []int{5, 3, 2, 2, 2}
	fmt.Println(isEqual(a, b))
}
