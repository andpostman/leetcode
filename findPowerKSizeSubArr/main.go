package main

import "fmt"

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 2, 5}, 3))
	fmt.Println(resultsArray([]int{2, 2, 2, 2, 2}, 4))
	fmt.Println(resultsArray([]int{3, 2, 3, 2, 3, 2}, 2))
}
func resultsArray(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	l := 0
	consecCnt := 1
	for r, num := range nums {
		if r > 0 && nums[r-1]+1 == num {
			consecCnt++
		}
		if r-l+1 > k {
			if nums[l]+1 == nums[l+1] {
				consecCnt--
			}
			l++
		}
		if r-l+1 == k {
			if consecCnt == k {
				res = append(res, num)
			} else {
				res = append(res, -1)
			}
		}
	}
	return res
}
