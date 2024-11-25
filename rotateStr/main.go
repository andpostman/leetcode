package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("bbbacddceeb", "ceebbbbacdd"))
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
