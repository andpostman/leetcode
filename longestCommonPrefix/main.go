package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix, maxPrefix string
	for _, char := range strs[0] {
		prefix += string(char)
		for i := 1; i < len(strs); {
			str := strs[i]
			if !strings.HasPrefix(str, prefix) {
				prefix = prefix[:len(prefix)-1]
			} else {
				i++
			}
			if prefix == "" || prefix == maxPrefix {
				return maxPrefix
			}
		}
		maxPrefix = prefix
	}
	return maxPrefix
}
