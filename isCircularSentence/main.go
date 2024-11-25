package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isCircularSentence("leetcode exercises sound delightful"))
	fmt.Println(isCircularSentence("eetcode"))
	fmt.Println(isCircularSentence("Leetcode is cool"))
}

func isCircularSentence2(sentence string) bool {
	words := strings.Fields(sentence)
	first := words[0]
	last := words[len(words)-1]
	if first[0] != last[len(last)-1] {
		return false
	}
	for i := 1; i < len(words); i++ {
		prev := words[i-1]
		cur := words[i]
		if prev[len(prev)-1] != cur[0] {
			return false
		}
	}
	return true
}

//or

func isCircularSentence(sentence string) bool {
	l := len(sentence)
	if sentence[0] != sentence[l-1] {
		return false
	}
	for i := 1; i < l-1; i++ {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return true
}
