package main

import (
	"fmt"
)

func main() {
	arr := []string{"abc", "ab", "bc", "b"}
	fmt.Println(sumPrefixScores(arr))
}

type Node struct {
	children map[string]*Node
	count int
}

func (n *Node) insert(word string) {
	for _, val := range word {
		next, exist := n.children[string(val)]
		if exist {
			n = next
			n.count++
		} else {
			newnode := &Node{children: make(map[string]*Node), count: 1}
			n.children[string(val)] = newnode
			n = newnode
		}
	}
}

func (n *Node) search(word string) int {
	counter := 0
	for _, val := range word {
		next, exist := n.children[string(val)]
		if exist {
			n = next
			counter+=n.count
		} else {
			return counter
		}
	}
	return counter
}

func sumPrefixScores(words []string) []int {
	root := &Node{children: make(map[string]*Node)}
	for _, word := range words {
		root.insert(word)
	}
	result := make([]int, len(words))
	for i, word := range words {
		result[i] = root.search(word)
	}
	return result
}

//func longestCommonPrefix(arr1 []int, arr2 []int) int {
//	root := &Node{children: make(map[string]*Node)}
//	for _, val := range arr1 {
//		root.insert(strconv.Itoa(val))
//	}
//
//	result := 0
//	for _, val := range arr2 {
//		result = max(root.search(strconv.Itoa(val)), result)
//	}
//	return result
//}
