package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minGroups([][]int{{441459, 446342}, {801308, 840640}, {871890, 963447}, {228525, 336985}, {807945, 946787}, {479815, 507766}, {693292, 944029}, {751962, 821744}}))
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pq := &MinHeap{}
	heap.Init(pq)
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if pq.Len() > 0 && (*pq)[0] < start {
			heap.Pop(pq)
		}
		heap.Push(pq, end)
	}
	return pq.Len()
}
