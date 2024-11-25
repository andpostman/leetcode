package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

type MaxHeap []int64

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	elems := &MaxHeap{}
	heap.Init(elems)
	for _, num := range nums {
		heap.Push(elems, int64(num))
	}
	var score int64
	for i := k; elems.Len() > 0 && i > 0; i-- {
		elem := heap.Pop(elems).(int64)
		score += elem
		heap.Push(elems, int64(math.Ceil(float64(elem)/3)))
	}
	return score
}
