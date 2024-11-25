package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, word := range words {
		str := strings.Trim(word, allowed)
		if str == "" {
			count++
		}
	}
	return count
}
