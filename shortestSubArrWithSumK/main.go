package main

import "math"

func main() {

}

func shortestSubarray(nums []int, k int) int {
	minSize := math.MaxInt32
	type pair struct {
		prefixSum int64
		endIdx    int
	}
	// Use a slice to implement deque functionality
	queue := make([]pair, 0)
	var curSum int64
	for r, num := range nums {
		curSum += int64(num)
		if curSum >= int64(k) {
			minSize = min(minSize, r+1)
		}
		// Find the minimum valid window ending at r
		for len(queue) > 0 && curSum-queue[0].prefixSum >= int64(k) {
			minSize = min(minSize, r-queue[0].endIdx)
			queue = queue[1:]
		}
		// Validate the monotonic deque
		for len(queue) > 0 && queue[len(queue)-1].prefixSum > curSum {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, pair{
			prefixSum: curSum,
			endIdx:    r,
		})
	}
	if minSize == math.MaxInt32 {
		return -1
	}
	return minSize
}
