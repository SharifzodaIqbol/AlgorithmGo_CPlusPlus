package main

import "fmt"

func add(a, b string) string {
	return a + b
}

func main() {
	a := "a"
	b := "b"
	fmt.Println(add(a, b))
}
