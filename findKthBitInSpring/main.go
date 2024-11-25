package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthBit(4, 11))
}

func findKthBit(n int, k int) byte {

	var s string
	for i := 0; i < n; i++ {
		if i == 0 {
			s = "0"
		} else {
			s = s + "1" + reverseInvert(s)
		}
	}
	return s[k-1]
}

func reverseInvert(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] == '0' {
			runes[i] = '1'
		} else {
			runes[i] = '0'
		}
	}
	for i := 0; i <= len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}
