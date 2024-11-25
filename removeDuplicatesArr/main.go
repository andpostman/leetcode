package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 2}))
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	unique := make(map[int]bool)
	//res := make([]int, 0, len(nums))
	idx := 0
	for _, num := range nums {
		if !unique[num] {
			unique[num] = true
			nums[idx] = num
			idx++
		}
	}
	return len(unique)
}

// second
func removeDuplicates2(nums []int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return len(nums[:left+1])
}
