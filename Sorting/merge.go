package main

import "fmt"

func merge(arr []int, left, mid, right int) {
	n1, n2 := mid-left+1, right-mid
	L := make([]int, n1)
	R := make([]int, n2)
	//fmt.Println("left =", left, "mid =", mid, "right=", right)
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	//fmt.Println("L=", L)
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}
	//fmt.Println("R=", R)
	i, j := 0, 0
	k := left
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}
func main() {
	arr := []int{38, 47, 43, 10}
	mergeSort(arr, 0, len(arr)-1) // 10 38 43 47
	fmt.Println(arr)
}
