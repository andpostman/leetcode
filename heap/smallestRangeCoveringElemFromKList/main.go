package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}))
}

type Element struct {
	value  int
	row    int
	column int
}

type MinHeap []Element

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].value < m[j].value
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Element))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	elements := &MinHeap{}
	maxValue := math.MinInt32
	heap.Init(elements)

	// Initialize heap with first element from each list
	for i, num := range nums {
		heap.Push(elements, Element{
			value:  num[0],
			row:    i,
			column: 0,
		})
		maxValue = max(maxValue, num[0])
	}

	rangeStart, rangeEnd := 0, math.MaxInt32

	for elements.Len() > 0 {
		minElem := heap.Pop(elements).(Element)
		// Update the smallest range
		if maxValue-minElem.value < rangeEnd-rangeStart {
			rangeStart = minElem.value
			rangeEnd = maxValue
		}
		// Move to the next element in the current list
		if minElem.column+1 < len(nums[minElem.row]) {
			nextValue := nums[minElem.row][minElem.column+1]
			heap.Push(elements, Element{
				value:  nextValue,
				row:    minElem.row,
				column: minElem.column + 1,
			})
			maxValue = max(maxValue, nextValue)
		} else {
			break // One list is exhausted
		}
	}
	return []int{rangeStart, rangeEnd}
}
