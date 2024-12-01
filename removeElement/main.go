package main

import "fmt"

func main() {
	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{3, 3}, 5))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 1{
        if nums[0] == val{
            return 0
        }
        return 1
    }
	notEqVal := 0
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left] == val {
			if nums[right] != val {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				notEqVal++
			} else {
				right--
			}
		} else {
			notEqVal++
			left++
		}
	}
	fmt.Println(nums)
	return notEqVal
}
