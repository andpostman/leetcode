package main

import "fmt"

func main() {
	fmt.Println(removeStones([][]int{
		{
			0, 0,
		},
		{
			0, 1,
		},
		{
			1, 0,
		},
		{
			1, 2,
		},
		{
			2, 1,
		},
		{
			2, 2,
		},
	}))
	fmt.Println(removeStones([][]int{
		{
			0, 0,
		},
		{
			0, 2,
		},
		{
			1, 1,
		},
		{
			2, 0,
		},
		{
			2, 2,
		},
	}))
	fmt.Println(removeStones([][]int{
		{
			0, 0,
		},
	}))
	fmt.Println(removeStones([][]int{
		{
			0, 1,
		},
		{
			1, 0,
		},
		{
			1, 1,
		},
	}))
	fmt.Println(removeStones([][]int{
		{
			0, 1,
		},
		{
			1, 0,
		},
	}))
	fmt.Println(removeStones([][]int{
		{
			2, 6,
		},
		{
			2, 0,
		},
		{
			4, 2,
		},
		{
			1, 0,
		},
		{
			5, 2,
		},
		{
			0, 2,
		},
		{
			6, 5,
		},
	}))
}

type Node struct {
	x, y int
}

type Graph struct {
	nodes map[Node][]Node
}

func (g *Graph) AddEdge(a, b Node) {
	g.nodes[a] = append(g.nodes[a], b)
	g.nodes[b] = append(g.nodes[b], a)
}

func Dfs(graph *Graph, start Node, visited map[Node]bool) []Node {
	if visited[start] {
		return nil
	}
	var order []Node
	var dfs func(node Node)
	dfs = func(node Node) {
		visited[node] = true
		order = append(order, node)

		for _, neighbor := range graph.nodes[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}
	dfs(start)
	return order
}

func removeStones(stones [][]int) int {
	graph := &Graph{nodes: make(map[Node][]Node)}
	var nodes []Node
	for i, stoneToCheck := range stones {
		mainNode := Node{stoneToCheck[0], stoneToCheck[1]}
		for j, stoneCoordinate := range stones {
			checkNode := Node{stoneCoordinate[0], stoneCoordinate[1]}
			if i != j && mainNode.x == checkNode.x || mainNode.y == checkNode.y {
				graph.AddEdge(mainNode, checkNode)
			}
		}
		nodes = append(nodes, mainNode)
	}
	sum := 0
	visited := make(map[Node]bool)
	for _, node := range nodes {
		dfs := Dfs(graph, node, visited)
		if dfs != nil {
			sum += len(dfs) - 1
		}

	}

	return sum
}
