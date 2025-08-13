package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	slice := []int{2, 3, 5, 4, 5, 3, 4}
	for _, value := range slice {
		myMap[value]++
	}
	for key, value := range myMap {
		if value == 1 {
			fmt.Println(key)
			return
		}
	}
	fmt.Println("Не существует уникальное число!")
}
