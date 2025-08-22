package main

import "fmt"

func isSubset(a []int, b []int) bool {
	Hash := make(map[int]struct{})
	for _, value := range a {
		Hash[value] = struct{}{}
	}
	for _, value := range b {
		if _, ok := Hash[value]; !ok {
			return false
		}
	}
	return true
}

func main() {
	a := []int{10, 5, 2, 23, 19}
	b := []int{19, 5, 3}
	fmt.Println(isSubset(a, b))

}
