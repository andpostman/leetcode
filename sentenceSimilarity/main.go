package main

import (
	"fmt"
	"strings"
)

func main() {
	res := areSentencesSimilar("My name is Haley", "My Haley")
	fmt.Println(res)
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	// Split sentences into words
	words1 := strings.Fields(sentence1)
	words2 := strings.Fields(sentence2)
	// Ensure words1 is the longer sentence
	if len(words2) > len(words1) {
		words1, words2 = words2, words1
	}

	start, end := 0, 0
	l1, l2 := len(words1), len(words2)
	// Compare from the start
	for start < l2 && words1[start] == words2[start] {
		start++
	}
	// Compare from the end
	for end < l2 && words1[l1-1-end] == words2[l2-1-end] {
		end++
	}
	// Check if the remaining unmatched part is in the middle
	return start+end >= l2
}
