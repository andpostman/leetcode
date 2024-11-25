package main

func main() {
	println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	println(maximumSubarraySum([]int{4, 4, 4}, 3))
	println(maximumSubarraySum([]int{5, 5, 3, 3, 1}, 3))
	println(maximumSubarraySum([]int{100, 1, 2, 100}, 2))
	println(maximumSubarraySum([]int{9, 9, 9, 1, 2, 3}, 3))
	println(maximumSubarraySum([]int{3, 3, 5, 5, 3, 3, 1, 1, 4, 2, 5, 1}, 4))
	println(maximumSubarraySum([]int{9, 18, 10, 13, 17, 9, 19, 2, 1, 18}, 5))
}

func maximumSubarraySum(nums []int, k int) int64 {
	count := make(map[int]int)
	var ans, curSum int64
	for r, l := 0, 0; r < len(nums); r++ {
		count[nums[r]]++
		curSum += int64(nums[r])
		for count[nums[r]] > 1 {
			count[nums[l]]--
			curSum -= int64(nums[l])
			l++
		}
		if r-l+1 > k {
			curSum -= int64(nums[l])
			count[nums[l]]--
			l++
		}
		if r-l+1 == k {
			ans = max(ans, curSum)
		}
	}
	return ans
}

// SECOND SOLUTION
func maximumSubarraySumSecond(nums []int, k int) int64 {
	var res, curSum int64
	count := make(map[int]int)

	for r, l := 0, 0; r < len(nums); r++ {
		curSum += int64(nums[r])
		count[nums[r]]++

		if r-l+1 > k {
			count[nums[l]]--
			if count[nums[l]] == 0 {
				delete(count, nums[l])
			}
			curSum -= int64(nums[l])
			l++
		}

		if len(count) == r-l+1 && r-l+1 == k {
			if curSum > res {
				res = curSum
			}
		}
	}

	return res
}
