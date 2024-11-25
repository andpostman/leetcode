package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Println(longestSquareStreak([]int{2, 3, 5, 6, 7}))
}

func longestSquareStreak(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}
	slices.Sort(nums)
	maxNum := 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		count := 0
		for seen[num] {
			count++
			num *= num
		}
		maxNum = max(count, maxNum)
		//With the constraints, the length of the longest square streak possible is 5.
		if maxNum == 5 {
			break
		}
	}
	if maxNum == 1 {
		return -1
	}
	return maxNum
}

// OR sec sol
func longestSquareStreak2(nums []int) int {
	mp := make(map[int]int)
	sort.Ints(nums)
	res := -1

	for _, num := range nums {
		sqrt := int(math.Sqrt(float64(num)))

		if sqrt*sqrt == num {
			if val, exists := mp[sqrt]; exists {
				mp[num] = val + 1
				if mp[num] > res {
					res = mp[num]
				}
			} else {
				mp[num] = 1
			}
		} else {
			mp[num] = 1
		}
	}

	return res
}
