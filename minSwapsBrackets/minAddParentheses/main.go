package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("()))(("))
}

func minAddToMakeValid(s string) int {
	stack := make([]rune, 0, len(s))
	for _, char := range s {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top == '(' && char == ')' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, char)
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack)
}
