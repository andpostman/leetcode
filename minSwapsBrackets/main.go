package main

import "fmt"

func main() {
	fmt.Println(minSwaps("]]][[["))
}

func minSwaps(s string) int {
	imbalance := 0
	swapCount := 0
	for _, char := range s {
		if char == ']' {
			if imbalance > 0 {
				imbalance--
			} else {
				imbalance++
				swapCount++
			}
		} else if char == '[' {
			imbalance++
		}
	}
	return swapCount
}
