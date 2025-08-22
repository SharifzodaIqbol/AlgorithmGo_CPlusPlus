package main

import "fmt"

func isDisjoint(a []int, b []int) bool {
	Hash := make(map[int]struct{})
	for _, value := range a {
		Hash[value] = struct{}{}
	}
	for _, value := range b {
		if _, ok := Hash[value]; ok {
			return false
		}
	}
	return true
}

func main() {
	a := []int{12, 34, 11, 9, 3}
	b := []int{7, 2, 1, 5}
	fmt.Println(isDisjoint(a, b))
}
