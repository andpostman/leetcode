package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(makeFancyString("aab"))
}

func makeFancyString(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n < 2 || !(result[n-1] == s[i] && result[n-2] == s[i]) {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// OR
func makeFancyString2(s string) string {
	curChar := s[0]
	curNum := 1
	builder := strings.Builder{}
	builder.WriteString(string(curChar))
	for i := 1; i < len(s); i++ {
		if next := s[i]; next == curChar {
			curNum++
		} else {
			curChar = next
			curNum = 1
		}
		if curNum < 3 {
			builder.WriteString(string(curChar))
		}
	}
	return builder.String()
}
