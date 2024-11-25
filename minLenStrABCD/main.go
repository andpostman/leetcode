package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(minLength("ABFCACDB"))
}

func minLength(s string) int {
	l := list.New()
	for _, char := range s {
		if l.Len() != 0 {
			top := l.Back().Value.(int32)
			if top == 'A' && char == 'B' || top == 'C' && char == 'D' {
				l.Remove(l.Back())
			} else {
				l.PushBack(char)
			}
		} else {
			l.PushBack(char)
		}
	}
	return l.Len()
}
