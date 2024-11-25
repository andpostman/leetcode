package main

import "fmt"

func main() {
	fmt.Println(minChanges("111101101110"))
}

func minChanges(s string) int {
	ans := 0
	for i := 1; i < len(s); i += 2 {
		if s[i-1] != s[i] {
			ans++
		}
	}
	return ans
}
