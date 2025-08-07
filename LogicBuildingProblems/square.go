package main

import (
	"fmt"
)

/*
O(sqrt(n)) Time and O(1) Space
*/
func floorSqrt(n int) int {
	result := 1
	for result*result <= n {
		result++
	}
	return result - 1
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(floorSqrt(n))
}

/*
-> FIRST METHOD
O(log(n)) Time and O(1) Space
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(int(math.Sqrt(float64(n))))
}
*/
/*
-> SECOND METHOD
O(1) Time and O(1) Space
func main() {
	var n int
	fmt.Scan(&n)
	SquareRoot := math.Pow(float64(n), 0.5)
	fmt.Println(math.Floor(SquareRoot))
}
*/
