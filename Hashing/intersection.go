package main

import "fmt"

func intersection(a, b []int) []int {
	Hash := make(map[int]struct{})
	slice := []int{}
	for _, value := range a {
		Hash[value] = struct{}{}
	}
	for _, value := range b {
		if _, ok := Hash[value]; ok {
			slice = append(slice, value)
			delete(Hash, value)
		}
		if len(Hash) == 0 {
			return slice
		}
	}
	return slice
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7, 8}
	fmt.Println(intersection(a, b))
}
