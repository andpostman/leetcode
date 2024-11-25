package main

import (
	"container/heap"
	"fmt"
)

func main() {
	res := maxProbability(10, [][]int{{0, 3}, {1, 7}, {1, 2}, {0, 9}}, []float64{0.31, 0.9, 0.86, 0.36}, 2, 3)
	fmt.Println(res)
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	g := NewGraph(n)
	for i, ints := range edges {
		g.AddEdge(ints[0], ints[1], succProb[i])
	}
	res := g.Dijkstra(start_node, end_node)
	if res == -1 {
		return 0
	}
	return res
}

type Graph struct {
	vertices [][]Vertex
}

func NewGraph(n int) *Graph {
	return &Graph{
		vertices: make([][]Vertex, n),
	}
}

func (g *Graph) AddEdge(from, to int, weight float64) {
	g.vertices[from] = append(g.vertices[from], Vertex{index: to, dist: weight})
	g.vertices[to] = append(g.vertices[to], Vertex{index: from, dist: weight}) // если граф неориентированный
}

type Vertex struct {
	index int
	dist  float64
}

type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Vertex))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) Dijkstra(from, to int) float64 {
	distances := make([]float64, len(g.vertices))
	for i := range distances {
		distances[i] = -1
	}
	distances[from] = 1

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, Vertex{from, 1})

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(Vertex)
		if v.index == to {
			return v.dist
		}

		for _, neighbor := range g.vertices[v.index] {
			newDist := v.dist * neighbor.dist
			if newDist > distances[neighbor.index] {
				distances[neighbor.index] = newDist
				heap.Push(&pq, Vertex{neighbor.index, newDist})
			}
		}
	}

	return -1
}
