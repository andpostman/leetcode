package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(treeQueries(root, []int{3, 2, 4, 8}))
}

func treeQueries(root *TreeNode, queries []int) []int {
	maxHeight := &MaxHeight{}
	// First traversal from left to right
	maxHeight.traversalLeftToRight(root, 0)
	// Reset current max height for second traversal
	maxHeight.currentMaxHeight = 0
	maxHeight.traversalRightToLeft(root, 0)
	// Process queries and build result
	results := make([]int, len(queries))
	for i, query := range queries {
		results[i] = maxHeight.maxHeightAfterRemoval[query]
	}
	return results
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MaxHeight struct {
	maxHeightAfterRemoval []int
	currentMaxHeight      int
}

func (m *MaxHeight) traversalLeftToRight(node *TreeNode, height int) {
	if node == nil {
		return
	}
	// Update len of arr (max arr 100001)
	if node.Val >= len(m.maxHeightAfterRemoval) {
		old := m.maxHeightAfterRemoval
		new := make([]int, node.Val*2)
		copy(new, old)
		m.maxHeightAfterRemoval = new
	}
	// Store the maximum height if this node were removed
	m.maxHeightAfterRemoval[node.Val] = m.currentMaxHeight
	// Update current maximum height
	m.currentMaxHeight = max(m.currentMaxHeight, height)
	// Traverse left subtree first, then right
	m.traversalLeftToRight(node.Left, height+1)
	m.traversalLeftToRight(node.Right, height+1)
}

func (m *MaxHeight) traversalRightToLeft(node *TreeNode, height int) {
	if node == nil {
		return
	}
	// Update maximum height if this node were removed
	m.maxHeightAfterRemoval[node.Val] = max(m.maxHeightAfterRemoval[node.Val], m.currentMaxHeight)
	// Update current maximum height
	m.currentMaxHeight = max(m.currentMaxHeight, height)
	// Traverse right subtree first, then left
	m.traversalRightToLeft(node.Right, height+1)
	m.traversalRightToLeft(node.Left, height+1)
}
