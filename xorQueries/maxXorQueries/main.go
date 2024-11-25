package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaximumXor([]int{2, 3, 4, 7}, 3))
}

func getMaximumXor(nums []int, maximumBit int) []int {
	//xorRes := int(math.Pow(2, float64(maximumBit))) - 1
	//THE SAME
	xorRes := (1 << maximumBit) - 1
	prefix := make([]int, len(nums))
	for i, num := range nums {
		xorRes ^= num
		prefix[len(prefix)-1-i] = xorRes
	}
	return prefix
}
