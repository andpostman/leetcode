package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 4}))
}

func longestSubarray(nums []int) int {
	maxVal := nums[0]
	maxLen := 1
	currentLen := 1
	for i := 1; i < len(nums); i++ {
		if maxVal < nums[i] {
			maxVal = nums[i]
			maxLen = 1
			currentLen = 1
		} else if maxVal == nums[i] {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 0
		}
	}
	return maxLen
}
