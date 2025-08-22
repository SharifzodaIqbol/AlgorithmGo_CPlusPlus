package main

import "fmt"

func maxDistance(arr []int) int {
	mx := 0
	Hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := Hash[arr[i]]; ok {
			mx = max(mx, i-Hash[arr[i]])
		} else {
			Hash[arr[i]] = i
		}
	}
	return mx
}
func main() {
	arr := []int{1, 2, 3, 6, 5, 4}
	fmt.Println(maxDistance(arr))

}
