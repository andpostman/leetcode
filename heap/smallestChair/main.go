package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
}

// IntHeap for storing available chairs as a min-heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap based on chair numbers
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// EventHeap for storing leaving events based on time
type EventHeap [][2]int

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // Min-heap based on leaving times
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)

	// Create a list of arrivals with friend index
	arrivals := make([][2]int, n)
	for i := 0; i < n; i++ {
		arrivals[i] = [2]int{times[i][0], i}
	}

	// Sort friends by arrival time
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i][0] < arrivals[j][0]
	})

	// Min-Heap to track available chairs
	availableChairs := &IntHeap{}
	heap.Init(availableChairs)
	for i := 0; i < n; i++ {
		heap.Push(availableChairs, i)
	}

	// Priority queue to track when chairs are freed (leavingQueue)
	leavingQueue := &EventHeap{}
	heap.Init(leavingQueue)

	// Iterate through each friend based on arrival
	for _, arrival := range arrivals {
		arrivalTime, friendIndex := arrival[0], arrival[1]

		// Free chairs that are vacated before the current arrival time
		for leavingQueue.Len() > 0 && (*leavingQueue)[0][0] <= arrivalTime {
			heap.Push(availableChairs, heap.Pop(leavingQueue).([2]int)[1])
		}

		// Assign the smallest available chair
		chair := heap.Pop(availableChairs).(int)

		// If this is the target friend, return their chair number
		if friendIndex == targetFriend {
			return chair
		}

		// Mark the chair as being used until the friend's leave time
		heap.Push(leavingQueue, [2]int{times[friendIndex][1], chair})
	}

	return -1 // Should never reach here
}

// Second solution

//type ChairHeap []int
//
//func (c ChairHeap) Len() int {
//	return len(c)
//}
//
//func (c ChairHeap) Less(i, j int) bool {
//	return c[i] < c[j]
//}
//
//func (c ChairHeap) Swap(i, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//func (c *ChairHeap) Push(x interface{}) {
//	*c = append(*c, x.(int))
//}
//
//func (c *ChairHeap) Pop() interface{} {
//	old := *c
//	n := len(old)
//	x := old[n-1]
//	*c = old[:n-1]
//	return x
//}
//
//type DepartureHeap [][2]int
//
//func (a DepartureHeap) Len() int {
//	return len(a)
//}
//
//func (a DepartureHeap) Less(i, j int) bool {
//	return a[i][0] < a[j][0]
//}
//
//func (a DepartureHeap) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//}
//
//func (a *DepartureHeap) Push(x interface{}) {
//	*a = append(*a, x.([2]int))
//}
//
//func (a *DepartureHeap) Pop() interface{} {
//	old := *a
//	n := len(old)
//	x := old[n-1]
//	*a = old[:n-1]
//	return x
//}
//
//func smallestChair2(times [][]int, targetFriend int) int {
//	n := len(times)
//	arrivals := make([][2]int, n)
//	availableChairs := &ChairHeap{}
//	heap.Init(availableChairs)
//	for i := 0; i < n; i++ {
//		arrivals[i] = [2]int{times[i][0], i}
//		heap.Push(availableChairs, i)
//	}
//	sort.Slice(arrivals, func(i, j int) bool {
//		return arrivals[i][0] < arrivals[j][0]
//	})
//	leavingQueue := &DepartureHeap{}
//	heap.Init(leavingQueue)
//	for _, arrival := range arrivals {
//		arrivalTime, friendIndex := arrival[0], arrival[1]
//		for leavingQueue.Len() > 0 && (*leavingQueue)[0][0] <= arrivalTime {
//			heap.Push(availableChairs, heap.Pop(leavingQueue).([2]int)[1])
//		}
//		chair := heap.Pop(availableChairs).(int)
//		if friendIndex == targetFriend {
//			return chair
//		}
//		heap.Push(leavingQueue, [2]int{times[friendIndex][1], chair})
//	}
//	return -1
//}
