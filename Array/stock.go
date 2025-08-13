package main

import "fmt"

func main() {
	slice := []int{100, 180, 260, 310, 40, 535, 695}
	stock := 0
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			stock += slice[i] - slice[i-1]
		}
	}
	fmt.Println(stock)
}
