package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var Hash = map[int]string{
		3: "Fizz",
		5: "Bazz",
	}
	keyHash := []int{3, 5}
	sliceStr := []string{}
	i := 1
	for i <= n {
		s := ""
		for _, value := range keyHash {
			if i%value == 0 {
				s += Hash[value]
			}
		}
		if len(s) == 0 {
			s += strconv.Itoa(i)
		}
		sliceStr = append(sliceStr, s)
		i++
	}
	return sliceStr
}

func main() {
	var n int = 20
	fmt.Println(fizzBuzz(n))
}
