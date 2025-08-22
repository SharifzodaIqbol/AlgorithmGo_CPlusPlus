package main

import "fmt"

func merge(arr1 []int, arr2 []int) {
	n1, n2 := len(arr1), len(arr2)
	i, j := 0, 0
	arr3 := make([]int, 0)
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}
	if i != n1 {
		arr3 = append(arr3, arr1[i:]...)
	} else if j != n2 {
		arr3 = append(arr3, arr2[j:]...)
	}
	copy(arr1, arr3[:n1])
	copy(arr2, arr3[n1:])

}

func main() {
	arr1 := []int{1, 3, 4, 5, 6}
	arr2 := []int{2, 4, 6, 8}
	merge(arr1, arr2)
	fmt.Println(arr1, arr2)
}
