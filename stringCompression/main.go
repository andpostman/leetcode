package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compressedString("abcde"))
	fmt.Println(compressedString("jd"))
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))
}

func compressedString(word string) string {
	prev := word[0]
	count := 1
	var comp strings.Builder
	// Initialize result builder with capacity to avoid reallocations
	comp.Grow(len(word) * 2)
	for i := 1; i < len(word); i++ {
		if word[i] != prev || count == 9 {
			comp.WriteByte(byte('0' + count))
			//comp.WriteString(fmt.Sprint(count))
			comp.WriteByte(prev)
			prev = word[i]
			count = 1
		} else {
			count++
		}
	}
	//comp.WriteString(fmt.Sprint(count) + string(prev))
	comp.WriteByte(byte('0' + count))
	comp.WriteByte(prev)
	return comp.String()
}
