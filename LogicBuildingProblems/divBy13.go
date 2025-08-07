package main

import "fmt"

func divBy13(str string) bool {
	rem := 0
	for i := 0; i < len(str); i++ {
		rem = (rem*10 + int(str[i]-'0')) % 13
	}
	fmt.Println(rem)
	return rem == 0
}
func main() {
	str := "2911285"
	fmt.Println(divBy13(str))

}
