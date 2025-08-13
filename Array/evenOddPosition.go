package main

import "fmt"

func main() {
	slice := []int{1, 2, 2, 1}
	for i := 1; i < len(slice); i++ {
		if i%2 == 0 {
			if slice[i] > slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		} else {
			if slice[i] < slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}
	}
	for _, value := range slice {
		fmt.Print(value, " ")
	}
	return
}
