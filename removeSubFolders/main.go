package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(removeSubfolders([]string{"/c", "/d/c/e"}))
	fmt.Println(removeSubfolders([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
}

func removeSubfolders(folder []string) []string {
	slices.Sort(folder)
	var result []string
	result = append(result, folder[0])
	for i := 1; i < len(folder); i++ {
		lastFolder := result[len(result)-1] + "/"
		curr := folder[i]
		if !strings.HasPrefix(curr, lastFolder) {
			result = append(result, curr)
		}
	}
	return result
}
