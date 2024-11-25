package main

import "fmt"

func main() {
	fmt.Println(minimumSteps("101"))
}

func minimumSteps(s string) int64 {
	var steps int64 = 0
	var count int64 = 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == '1' {
			steps += count
		} else {
			count++
		}
	}
	return steps
}
