package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 10, 100}
	arr2 := []int{10, 1000}
	fmt.Println(longestCommonPrefix(arr1, arr2))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	m := make(map[string]bool)
	for _, elem := range arr1 {
		s := fmt.Sprint(elem)
		for i := 1; i <= len(s); i++ {
			m[s[:i]] = true
		}
	}
	max := 0
	for _, elem := range arr2 {
		s := fmt.Sprint(elem)
		for i := len(s); i >= 1; i-- {
			if m[s[:i]] {
				if max < i {
					max = i
					break
				}
			}
		}
	}
	return max
}
