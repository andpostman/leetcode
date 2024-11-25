package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

func diffWaysToCompute(expression string) []int {
	memory := make(map[string][]int)
	return divideAndConquer(expression, memory)
}

// Алгоритм разделяй и властвуй с memorization(map)
func divideAndConquer(expression string, memory map[string][]int) []int {
	if expressionResult, present := memory[expression]; present {
		return expressionResult
	}
	var result []int
	for i, char := range expression {
		if char == '-' || char == '*' || char == '+' {
			left := divideAndConquer(expression[:i], memory)
			right := divideAndConquer(expression[i+1:], memory)
			for _, l := range left {
				for _, r := range right {
					switch char {
					case '-':
						result = append(result, l-r)
					case '+':
						result = append(result, l+r)
					case '*':
						result = append(result, l*r)
					}
				}
			}

		}
	}
	if len(result) == 0 {
		number, _ := strconv.Atoi(expression)
		result = append(result, number)
	}
	memory[expression] = result
	return result
}
