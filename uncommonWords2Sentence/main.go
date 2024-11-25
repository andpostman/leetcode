package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q", uncommonFromSentences("apple apple", "banana"))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	commonSentence := s1 + " " + s2
	letters := strings.Fields(commonSentence)
	var res []string
	for _, letter := range letters {
		count := 0
		for _, let := range letters {
			if letter == let {
				count++
			}
			if count > 1 {
				break
			}
		}
		if count == 1 {
			res = append(res, letter)
		}
	}
	return res
}

//func uncommonFromSentences(s1 string, s2 string) []string {
//	commonSentence := s1 + " " + s2
//	letters := strings.Fields(commonSentence)
//	set := make(map[string]int)
//	for _, letter := range letters {
//		set[letter]++
//	}
//	var res []string
//	for key, value := range set {
//		if value == 1 {
//			res = append(res, key)
//		}
//	}
//	return res
//}
