package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	num := strings.Split(fmt.Sprint(x), "")
	left := 0
	right := len(num) - 1
	for left < right {
		if num[left] != num[right] {
			return false
		}
		left++
		right--
	}
	return true
}
