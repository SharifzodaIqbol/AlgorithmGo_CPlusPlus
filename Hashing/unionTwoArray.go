package main

import "fmt"

func union(a, b []int) []int {
	Hash := make(map[int]struct{})
	slice := []int{}
	for _, value := range a {
		Hash[value] = struct{}{}
	}
	for _, value := range b {
		if _, ok := Hash[value]; !ok {
			Hash[value] = struct{}{}
		}
	}
	for key, _ := range Hash {
		slice = append(slice, key)
	}
	return slice
}
func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Println(union(a, b))
}
