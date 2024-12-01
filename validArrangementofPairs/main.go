package main

import "fmt"

func main() {
	fmt.Println(validArrangement([][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}))
}

func validArrangement(pairs [][]int) [][]int {
	// Step 1: Build adjacency list and track in-out degrees
	adjacencyList := make(map[int][]int)
	inOutDegree := make(map[int]int)
	// Build graph and count in/out degrees
	for _, pair := range pairs {
		start, end := pair[0], pair[1]
		adjacencyList[start] = append(adjacencyList[start], end)
		inOutDegree[start]++ // out-degree
		inOutDegree[end]--   // in-degree
	}
	// Step 2: Find starting node
	startNode := pairs[0][0]
	for node, degree := range inOutDegree {
		if degree == 1 {
			startNode = node
			break
		}
	}
	// Step 3: Hierholzer's Algorithm for Eulerian path
	// Use slice as stack for path
	var path []int
	nodeStack := []int{startNode}
	for len(nodeStack) > 0 {
		lastIdx := len(nodeStack) - 1
		curr := nodeStack[lastIdx]
		neighbors := adjacencyList[curr]
		if len(neighbors) == 0 {
			// No more neighbors, add to path and pop stack
			path = append(path, curr)
			nodeStack = nodeStack[:lastIdx]
		} else {
			// Push next neighbor onto stack and remove the edge
			lastNeighborIdx := len(neighbors) - 1
			nextNode := neighbors[lastNeighborIdx]
			nodeStack = append(nodeStack, nextNode)
			adjacencyList[curr] = neighbors[:lastNeighborIdx]
		}
	}
	// Step 4: Reconstruct the arrangement
	// Build final arrangement
	arrangement := make([][]int, 0, len(path)-1)
	for i := len(path) - 1; i > 0; i-- {
		arrangement = append(arrangement, []int{path[i], path[i-1]})
	}
	return arrangement
}
