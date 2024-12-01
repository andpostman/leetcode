package main

import (
	"container/heap"
	"container/list"
	"math"
)

func main() {
	println(minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}))
}

type Edge struct {
	idx    [2]int
	weight int
}

type PriorityQueue []Edge

type Graph struct {
	edges [][][]Edge
}

// Dijkstra solution
func minimumObstacles(grid [][]int) int {
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])
	graph := NewGraph(m)
	for i := range grid {
		graph.edges[i] = make([][]Edge, n)
		for j := range grid[i] {
			for _, ints := range direction {
				newI := i + ints[0]
				newJ := j + ints[1]
				if newI >= 0 && newI < m && newJ >= 0 && newJ < n {
					graph.addEdge([2]int{i, j}, [2]int{newI, newJ}, grid[newI][newJ])
				}
			}
		}
	}
	return graph.Dijkstra([2]int{0, 0}, [2]int{m - 1, n - 1})
}

func NewGraph(n int) *Graph {
	return &Graph{edges: make([][][]Edge, n)}
}

func (g *Graph) Dijkstra(from, to [2]int) int {
	m := len(g.edges)
	n := len(g.edges[0])
	distances := make([][]int, m)
	for i := 0; i < m; i++ {
		distances[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distances[i][j] = math.MaxInt32
		}
	}
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, Edge{from, 0})
	for pq.Len() > 0 {
		edge := heap.Pop(&pq).(Edge)
		if edge.idx == to {
			return edge.weight
		}
		for _, neighbor := range g.edges[edge.idx[0]][edge.idx[1]] {
			curr := edge.weight + neighbor.weight
			if curr < distances[neighbor.idx[0]][neighbor.idx[1]] {
				distances[neighbor.idx[0]][neighbor.idx[1]] = curr
				heap.Push(&pq, Edge{
					idx:    neighbor.idx,
					weight: curr,
				})
			}
		}
	}
	return -1
}

func (g *Graph) addEdge(from, to [2]int, weight int) {
	g.edges[from[0]][from[1]] = append(g.edges[from[0]][from[1]], Edge{
		idx:    to,
		weight: weight,
	})
}

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
	*pq = append(*pq, x.(Edge))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// BFS solution
type PointState struct {
	x      int
	y      int
	weight int
}

func minimumObstaclesSec(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := list.New()
	queue.PushBack(PointState{0, 0, 0})
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	directions := []PointState{
		{0, 1, 0},
		{1, 0, 0},
		{0, -1, 0},
		{-1, 0, 0},
	}
	minObstacle := math.MaxInt32
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(PointState)
		for _, direction := range directions {
			nextX := curr.x + direction.x
			nextY := curr.y + direction.y
			weight := curr.weight
			if nextX == rows-1 && nextY == cols-1 {
				minObstacle = min(minObstacle, weight)
				continue
			}
			if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols && !visited[nextX][nextY] {
				visited[nextX][nextY] = true
				if grid[nextX][nextY] == 1 {
					weight++
					nextState := PointState{nextX, nextY, weight}
					queue.PushBack(nextState)
				} else {
					nextState := PointState{nextX, nextY, weight}
					queue.PushFront(nextState)
				}
			}
		}
	}
	return minObstacle
}
