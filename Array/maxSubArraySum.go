package main

import "fmt"

func main() {
	arr := []int{2, 3, -8, 7, -1, 2, 3}
	maxEnding := arr[0]
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		maxEnding = max(arr[i], arr[i]+maxEnding)
		result = max(result, maxEnding)
	}
	fmt.Println(result)
}
