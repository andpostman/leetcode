package main

import "math"

func main() {

}

func minimumSubarrayLength(nums []int, k int) int {
	result := 220000
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			return 1
		}
		n = n | nums[i]
		if n >= k {
			t := nums[i]
			j := i
			for t < k {
				j--
				t = t | nums[j]
			}
			if i-j+1 < result {
				result = i - j + 1
			}
			j++
			n = nums[j]
			i = j
		}
	}
	if result < 220000 {
		return result
	}
	return -1
}

//there are two utility method i used :-
//
//updateBits => this method gives me the number of set bit at each position in a given number stored in an array say count.
//
//getVal(int[]count)==> this method return me the equilvalent number for the array count that we produced in method of step1 updateBits.
//
//now we will solve our problem using Sliding window technique :-
//we will be having two variable start=0 and end=0
//
//using a while loop (end < arrayLength) ===>
//
//we would call the bitCount method this will update the set positions int count array for the nums[end] integer.
//
//right after we would check getVal(count) if this value is >=k that means at least a valid sub array we found at end index, so we will store the ans = math.min(ans, end-start+1)
//
//BUT we are looking for smallest sub array, so lets look for further smaller sub array
//
//so lets shrink the window, how? ===> by doing start++
//so we have to call bitCount to remove if some set bits are removed that we sitting at start index
//
//===> updateBit(count, nums[start], -1) ===> why -1? ====> because this time we are calling for decrease set bit at start index position.
//
//and keep on repeating the same, finally we would have smallest ans.
//
//BUT there might be a case that no sub array existing there
//in this case as initially our ans was initialized with Integer.MAX_VALUE, ===> so check if its value is still MAX, then return -1 else return ans.

func minimumSubarrayLengthSecond(nums []int, k int) int {
	count := make([]int, 32)
	start, end, res := 0, 0, math.MaxInt32
	for end < len(nums) {
		updateBits(count, nums[end], 1)
		for start <= end && getVal(count) >= k {
			res = min(res, end-start+1)
			updateBits(count, nums[start], -1)
			start++
		}
		end++
	}
	if res < math.MaxInt32 {
		return res
	}
	return -1
}

func updateBits(count []int, num, val int) {
	for i := 0; i < 32; i++ {
		if ((num >> i) & 1) == 1 {
			count[i] += val
		}
	}
}

func getVal(count []int) int {
	val := 0
	for i := 0; i < 32; i++ {
		if count[i] > 0 {
			val = val | (1 << i)
		}
	}
	return val
}
