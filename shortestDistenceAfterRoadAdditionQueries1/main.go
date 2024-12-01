package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestDistanceAfterQueriesSec(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
	fmt.Println(shortestDistanceAfterQueriesSec(4, [][]int{{0, 3}, {0, 2}}))
}

// THIRST SOLUTION DIJKSTRA
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := NewGraph(n)
	ans := make([]int, len(queries))
	for i, query := range queries {
		graph.addEdge(query[0], query[1], 1)
		ans[i] = graph.Dijkstra(0, n-1)
	}
	return ans
}

type Vertex struct {
	idx    int
	weight int
}

type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Vertex))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

type Graph struct {
	vertices [][]Vertex
}

func NewGraph(n int) *Graph {
	graph := &Graph{vertices: make([][]Vertex, n)}
	for i := 1; i < n; i++ {
		graph.addEdge(i-1, i, 1)
	}
	return graph
}

func (g *Graph) addEdge(from, to int, weight int) {
	g.vertices[from] = append(g.vertices[from], Vertex{idx: to, weight: weight})
}

func (g *Graph) Dijkstra(from, to int) int {
	distances := make([]int, len(g.vertices))
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, Vertex{from, 0})
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(Vertex)
		if v.idx == to {
			return v.weight
		}
		for _, neighbor := range g.vertices[v.idx] {
			currWeight := v.weight + neighbor.weight
			if currWeight < distances[neighbor.idx] {
				distances[neighbor.idx] = currWeight
				heap.Push(&pq, Vertex{neighbor.idx, currWeight})
			}
		}
	}
	return -1
}

// SECOND SOLUTION BFS
func shortestDistanceAfterQueriesSec(n int, queries [][]int) []int {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = []int{i + 1}
	}

	shortestPath := func() int {
		q := [][]int{{0, 0}} // node, length
		visit := make(map[int]bool)
		visit[0] = true

		for len(q) > 0 {
			cur, length := q[0][0], q[0][1]
			q = q[1:]

			if cur == n-1 {
				return length
			}

			for _, nei := range adj[cur] {
				if !visit[nei] {
					q = append(q, []int{nei, length + 1})
					visit[nei] = true
				}
			}
		}
		return -1
	}

	res := make([]int, 0)
	for _, query := range queries {
		src, dst := query[0], query[1]
		adj[src] = append(adj[src], dst)
		res = append(res, shortestPath())
	}
	return res
}

// THIRD SOLUTION DFS
func updateDistances(graph [][]int, current int, distances []int) {
	newDist := distances[current] + 1

	for _, neighbor := range graph[current] {
		if distances[neighbor] <= newDist {
			continue
		}

		distances[neighbor] = newDist
		updateDistances(graph, neighbor, distances)
	}
}

func shortestDistanceAfterQueriesThird(n int, queries [][]int) []int {
	distances := make([]int, n)
	for i := range distances {
		distances[i] = n - 1 - i
	}

	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i+1] = append(graph[i+1], i)
	}

	answer := make([]int, 0, len(queries))

	for _, query := range queries {
		source, target := query[0], query[1]
		graph[target] = append(graph[target], source)

		if distances[target]+1 < distances[source] {
			distances[source] = distances[target] + 1
		}

		updateDistances(graph, source, distances)
		answer = append(answer, distances[0])
	}

	return answer
}
