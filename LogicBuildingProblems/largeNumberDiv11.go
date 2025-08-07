package main

import "fmt"

func checkDivisible11(str string) bool {
	size := len(str)
	bytes := []byte(str)
	sumOdd, sumEvent := 0, 0
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			sumEvent += int(bytes[i] - byte('0'))
		} else {
			sumOdd += int(bytes[i] - byte('0'))
		}
	}
	if (sumEvent-sumOdd)%11 == 0 {
		return true
	}
	return false
}
func main() {
	var str string
	str = "7695"
	fmt.Println(checkDivisible11(str))
}
