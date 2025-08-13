package main

import "fmt"

func kDistance(slice []int, k int) bool {
	mp := make(map[int]int)
	for i, value := range slice {
		if _, ok := mp[value]; ok && i-mp[value] <= k {
			return true
		}
		mp[value] = i
	}
	return false
}
func main() {
	slice := []int{1, 2, 3, 4, 1, 2, 3, 4}
	k := 3
	fmt.Println(kDistance(slice, k))
}
