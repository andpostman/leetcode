package main

import (
	"fmt"
	"math/bits"
)

func main() {

	fmt.Println(minBitFlips(243, 640))
}

//
//func minBitFlips(start int, goal int) int {
//	result := start ^ goal
//	resultBits := strconv.FormatInt(int64(result), 2)
//	count := strings.Count(resultBits, "1")
//	return count
//}

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
