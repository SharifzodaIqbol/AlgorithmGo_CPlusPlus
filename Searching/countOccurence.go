package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 2, 1, 2, 3, 3}
	mp := make(map[int]int)
	k := 3
	k = len(arr) / k
	for _, value := range arr {
		mp[value]++
	}
	for key, value := range mp {
		if value > k {
			fmt.Printf("%d ", key)
		}
	}
}
