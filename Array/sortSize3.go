package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{5, 8, 2, 2, 2, 2, 2, 2, 2, 2, 10}

	mn := arr[0]

	mx := math.MaxInt32
	store := mn
	for _, value := range arr {
		if value == mn {
			continue
		} else if value < mn {
			mn = value
		} else if mx > value {
			mx = value
			store = mn
		} else if value > mx {
			fmt.Println(store, mx, value)
			return
		}
	}

}
