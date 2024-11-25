package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minimizedMaximumSec(6, []int{11, 6}))
	fmt.Println(minimizedMaximumSec(2, []int{5, 7}))
	fmt.Println(minimizedMaximumSec(7, []int{15, 10, 10}))
}

func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, slices.Max(quantities)
	result := 0
	for left <= right {
		mid := (right + left) / 2
		gap := 0
		for _, quantity := range quantities {
			reminder := int(math.Mod(float64(quantity), float64(mid)))
			quotient := quantity / mid
			if reminder > 0 {
				quotient++
			}
			gap += quotient
		}
		if gap <= n {
			right = mid - 1
			result = mid
		} else {
			left = mid + 1
		}
	}
	return result
}

// Second SOLUTION
func minimizedMaximumSec(n int, quantities []int) int {
	canDistribute := func(x int) bool {
		stores := 0
		for _, quantity := range quantities {
			stores += (quantity + x - 1) / x
		}
		return stores <= n
	}
	left, right := 1, slices.Max(quantities)
	result := 0
	for left <= right {
		x := (right + left) / 2
		if canDistribute(x) {
			result = x
			right = x - 1
		} else {
			left = x + 1
		}
	}
	return result
}

//This problem is not difficult if you solved similar problems with binary search.
//The main idea in such problems:
//
//find the range of possible answers (in this case 1..max_element(quantities)),
//guess the answer (number of shops in this task), at the middle of the range,
//try to check if this answer is possible (possible to distribute quantities to the shops),
//if not possible - move one range border to the middle or, if possible, another one,
//repeat.
