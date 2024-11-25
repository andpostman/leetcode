package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dividePlayers([]int{1, 1, 2, 3}))
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	checked := make(map[int64]bool)
	var sum int64 = 0
	for i, j := 0, len(skill)-1; i < j; i, j = i+1, j-1 {
		team := int64(skill[i] + skill[j])
		if !checked[team] {
			checked[team] = true
		}
		if len(checked) > 1 {
			return -1
		}
		sum += int64(skill[i] * skill[j])
	}
	return sum
}
