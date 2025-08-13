package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	sum := 0
	result := 0
	n := len(slice)
	for i, value := range slice {
		slice[i] = value * (n - i)
		sum += slice[i]
	}
	for _, value := range slice {
		result += sum
		sum -= value
	}
	fmt.Println(result)

}
