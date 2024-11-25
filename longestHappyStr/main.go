package main

import (
	"container/heap"
	"fmt"
	"strings"
)

//0 3 8
// ccbccbccbcc
//0 2 5
//0 5 8
// ccbbccbc
// ccbbcc
// ccbbccbbccbcc

func main() {
	fmt.Println(longestDiverseString(0, 8, 11))
}

type HappyStr struct {
	value rune
	count int
}

type MaxHeap []HappyStr

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].count > m[j].count
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(HappyStr))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	if a > 0 {
		heap.Push(maxHeap, HappyStr{value: 'a', count: a})
	}
	if b > 0 {
		heap.Push(maxHeap, HappyStr{value: 'b', count: b})
	}
	if c > 0 {
		heap.Push(maxHeap, HappyStr{value: 'c', count: c})
	}
	resultBuilder := strings.Builder{}
	for maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).(HappyStr)
		l := resultBuilder.Len()
		if l >= 2 {
			if rune(resultBuilder.String()[l-1]) != top.value ||
				rune(resultBuilder.String()[l-2]) != top.value {
				resultBuilder.WriteRune(top.value)
				top.count--
				if top.count > 0 {
					heap.Push(maxHeap, top)
				}
			} else {
				if maxHeap.Len() == 0 {
					break
				}
				next := heap.Pop(maxHeap).(HappyStr)
				resultBuilder.WriteRune(next.value)
				resultBuilder.WriteRune(top.value)
				next.count--
				top.count--
				if top.count > 0 {
					heap.Push(maxHeap, top)
				}
				if next.count > 0 {
					heap.Push(maxHeap, next)
				}
			}
		} else {
			resultBuilder.WriteRune(top.value)
			top.count--
			if top.count > 0 {
				heap.Push(maxHeap, top)
			}
		}
	}
	return resultBuilder.String()
}
