package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{10, 1000}}, []int{5}))
}

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] > items[j][1]
	})
	//fmt.Println(items)
	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		prevLen := len(ans)
		for _, item := range items {
			if query >= item[0] {
				ans = append(ans, item[1])
				break
			}
		}
		if prevLen == len(ans) {
			ans = append(ans, 0)
		}
	}
	return ans
}
