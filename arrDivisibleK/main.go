package main

import (
	"fmt"
)

func main() {
	fmt.Println(canArrange([]int{-2, -4, -9, -3, -1, -17}, 6))
}

func canArrange(arr []int, k int) bool {
	// Frequency array to store the count of remainders
	remainderFreq := make([]int, k)

	// Step 1: Calculate the remainder for each element and store the frequency
	for _, num := range arr {
		remainder := ((num % k) + k) % k // Ensure non-negative remainder
		remainderFreq[remainder]++
	}
	// Step 2: Check if the pairing condition holds
	for i := 0; i <= k/2; i++ {
		if i == 0 {
			// Elements with remainder 0 must pair among themselves
			if remainderFreq[i]%2 != 0 {
				return false
			}
		} else {
			// Remainder i must pair with remainder k-i
			if remainderFreq[i] != remainderFreq[k-i] {
				return false
			}
		}
	}
	return true
}

//func canArrange(arr []int, k int) bool {
//	var freq []int
//	checked := make(map[int]bool)
//	for _, el := range arr {
//		frequencies := ((el % k) + k) % k
//		freq = append(freq, frequencies)
//	}
//	sort.Ints(freq)
//	result := len(arr) / 2
//	addPair := 1
//	for result != 0 && addPair != 0 {
//		addPair = 0
//		for i, j := 0, len(freq)-1; i < j; {
//			if !checked[i] && !checked[j] {
//				sum := freq[i] + freq[j]
//				if sum == 0 || sum == k {
//					result--
//					checked[i] = true
//					checked[j] = true
//					addPair++
//					j--
//					i++
//				} else {
//					if sum > k {
//						j--
//					} else {
//						i++
//					}
//				}
//			} else {
//				if checked[i] {
//					i++
//				}
//				if checked[j] {
//					j--
//				}
//			}
//		}
//	}
//	if result == 0 {
//		return true
//	}
//
//	return false
//}
