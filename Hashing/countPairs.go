package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func countPairs(arr []int, k int) int {
	Hash := make(map[int]int)
	cnt := 0
	for _, value := range arr {
		if _, ok := Hash[value-k]; ok {
			cnt += Hash[value-k] // 1 + 2
		}
		if _, ok := Hash[value+k]; ok {
			cnt += Hash[value+k] // 1
		}
		Hash[value]++
	}
	/* k = 3
	1 1
	4 1 -> 1
	1 2 -> 1
	4 2 -> 2
	5 1 ->
	*/
	return cnt
}

func main() {
	arr := []int{1, 4, 1, 4, 5}
	k := 3
	fmt.Println(countPairs(arr, k))
}
