package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	maxStr := string(s[0])

	for i := 0; i < len(s)-1; i++ {
		odd := expandFromCenter(s, i, i)
		even := expandFromCenter(s, i, i+1)

		if len(even) > len(maxStr) {
			maxStr = even
		}

		if len(odd) > len(maxStr) {
			maxStr = odd
		}
	}

	return maxStr
}

func expandFromCenter(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return str[left+1 : right]
}

