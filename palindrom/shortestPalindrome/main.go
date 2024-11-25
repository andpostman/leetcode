package main

import "fmt"

func shortestPalindrome(s string) string {
	pattern := s + "#" + reverse(s)
	res := computePrefixSuffix(pattern)
	fmt.Println(res)
	return reverse(s[res[len(res)-1]:]) + s
}

// вычисляет префикс-функцию для строки-образца
func computePrefixSuffix(pattern string) []int {
	prefix := make([]int, len(pattern))
	index := 0
	for i := 1; i < len(pattern); i++ {
		for index > 0 && pattern[i] != pattern[index] {
			index = prefix[index-1]
		}
		if pattern[i] == pattern[index] {
			index++
		}
		prefix[i] = index
	}
	return prefix
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	text := "babad"
	fmt.Println(shortestPalindrome(text))
	text = "cbbd"
	fmt.Println(shortestPalindrome(text))

}
