package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	res := twoSum([]int{-6, 1, 0, 4}, -5)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var elements Elements
	for i, num := range nums {
		elements = append(elements, Element{i, num})
	}
	sort.Sort(elements)
	l, r := 0, len(nums)-1
	for left, right := l, r; left < right; {
		leftEl := elements[left].num
		rightEl := elements[right].num
		sum := elements[left].num + elements[right].num
		if sum == target {
			return []int{elements[left].idx, elements[right].idx}
		} else if leftEl < 0 && rightEl < 0 {
			if math.Abs(float64(sum)) < math.Abs(float64(target)) {
				right--
			} else {
				left++
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

type Element struct {
	idx int
	num int
}

type Elements []Element

func (e Elements) Len() int {
	return len(e)
}

func (e Elements) Less(i int, j int) bool {
	return e[i].num < e[j].num
}

func (e Elements) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}
