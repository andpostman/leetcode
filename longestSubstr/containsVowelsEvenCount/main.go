package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
}

func findTheLongestSubstring(s string) int {
	mapy := [32]int{}

	for i := range mapy {
		mapy[i] = -2
	}

	mapy[0] = -1
	maxLen := 0
	mask := 0

	for i, char := range s {
		switch char {
		case 'a':
			mask ^= 1
		case 'e':
			mask ^= 2
		case 'i':
			mask ^= 4
		case 'o':
			mask ^= 8
		case 'u':
			mask ^= 16
		}

		if mapy[mask] == -2 {
			mapy[mask] = i
		} else {
			length := i - mapy[mask]
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
