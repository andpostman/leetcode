package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	parentheses := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack []int32
	for _, r := range s {
		if _, ok := parentheses[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || parentheses[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
