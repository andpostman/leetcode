package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{999999, 999999998, 999999997}))
	fmt.Println(largestNumber([]int{1, 11, 111, 1112}))
}

func largestNumber(nums []int) string {
	numStr := make([]string, len(nums))
	for i, num := range nums {
		numStr[i] = strconv.Itoa(num)
	}
	res := mergeSort(numStr)
	if res[0] == "0" {
		return fmt.Sprint(res[0])
	}
	result := strings.Join(res, "")
	return result
}

func mergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []string) []string {
	var result []string
	for len(left) > 0 && len(right) > 0 {
		a := left[0] + right[0]
		b := right[0] + left[0]
		if a > b {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}
