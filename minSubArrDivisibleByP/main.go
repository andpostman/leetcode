package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))
}

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	reminder := sum % p
	if reminder == 0{
		return 0
	}
	checked := map[int]int{0: -1}
	minLen := len(nums)
	prefixSum := 0
	for i, num := range nums {
		prefixSum += num
		currMod := prefixSum % p
		targetMod := (currMod - reminder + p) % p

		if idx, ok := checked[targetMod]; ok {
			minLen = min(minLen, i-idx)
		} 
		checked[currMod] = i
	}

	if minLen == len(nums) {
		return -1
	}
	return minLen
}
