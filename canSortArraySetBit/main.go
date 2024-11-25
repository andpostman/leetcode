package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	fmt.Println(canSortArray([]int{3, 16, 8, 4, 2}))
}

func canSortArray(nums []int) bool {
	if slices.IsSorted(nums) {
		return true
	}
	prevKey := -1
	prevMax := -1
	currMax := -1
	for _, num := range nums {
		key := bits.OnesCount(uint(num))
		if prevKey != key {
			prevMax = currMax
			currMax = num
			prevKey = key
		} else {
			currMax = max(currMax, num)
		}
		if num < prevMax {
			return false
		}
	}
	return true
}
