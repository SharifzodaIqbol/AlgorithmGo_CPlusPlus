package main

import "fmt"

func check(freq []int) bool {
	size := len(freq)
	for i := 0; i < size; i++ {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}

func arePremutations(txt, pat string) bool {
	n := len(txt)
	m := len(pat)
	if m > n {
		return false
	}
	freq := [26]int{}
	for i := 0; i < m; i++ {
		freq[txt[i]-'a']++
		freq[pat[i]-'a']--
	}
	if check(freq[:]) {
		return true
	}
	for i := m; i < n; i++ {
		freq[txt[i]-'a']++
		freq[txt[i-m]-'a']--
		if check(freq[:]) {
			return true
		}
	}
	return false
}

func main() {
	txt := "programming"
	pat := "rain"
	fmt.Println(arePremutations(txt, pat))
}
