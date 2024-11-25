package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}

//func arrayRankTransform(arr []int) []int {
//	temp := make([]int, len(arr), cap(arr))
//	rankedArr := make([]int, len(arr), cap(arr))
//
//	copy(temp, arr)
//	sort.Ints(temp)
//	unique := make([]int, 0, cap(temp))
//	if len(temp) > 0 {
//		unique = append(unique, temp[0])
//	}
//	for i := 1; i < len(temp); i++ {
//		if temp[i-1] != temp[i] {
//			unique = append(unique, temp[i])
//		}
//	}
//
//	for i, num := range arr {
//		for j, sortNum := range unique {
//			if num == sortNum {
//				rankedArr[i] = j + 1
//				break
//			}
//		}
//	}
//
//	return rankedArr
//}

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	// Step 1: Create a sorted copy of the array
	temp := make([]int, len(arr), cap(arr))
	copy(temp, arr)
	sort.Ints(temp)
	// Step 2: Create a map to assign ranks
	ranked := make(map[int]int)
	rank := 1
	// Step 3: Assign ranks to sorted elements
	for _, num := range temp {
		if _, ok := ranked[num]; !ok {
			ranked[num] = rank
			rank++
		}
	}
	// Step 4: Replace each element in the original array with its rank
	for i, num := range arr {
		arr[i] = ranked[num]
	}
	return arr
}
