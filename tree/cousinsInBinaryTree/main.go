package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  9,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(replaceValueInTree(root))
}

type NodeInfo struct {
	root  *TreeNode
	level int
	sum   int
}

func replaceValueInTree(root *TreeNode) *TreeNode {

	queue := list.New()
	memorry := list.New()
	level := 0
	var levelSums []int
	levelSums = append(levelSums, root.Val)

	queue.PushBack(&NodeInfo{root, level, root.Val})
	memorry.PushBack(&NodeInfo{root, level, root.Val})

	for queue.Len() > 0 {

		var levelSum int

		pending := queue.Len()

		for i := 0; i < pending; i++ {

			nodeInfo := queue.Remove(queue.Front()).(*NodeInfo)
			node := nodeInfo.root
			sum := 0

			if node.Left != nil {
				sum += node.Left.Val
			}
			if node.Right != nil {
				sum += node.Right.Val
			}
			if node.Left != nil {
				queue.PushBack(&NodeInfo{node.Left, level + 1, sum})
				memorry.PushBack(&NodeInfo{node.Left, level + 1, sum})
			}
			if node.Right != nil {
				queue.PushBack(&NodeInfo{node.Right, level + 1, sum})
				memorry.PushBack(&NodeInfo{node.Right, level + 1, sum})
			}
			levelSum += sum
		}
		levelSums = append(levelSums, levelSum)
		level++
	}

	for memorry.Len() > 0 {
		nodeInfo := memorry.Remove(memorry.Front()).(*NodeInfo)
		node := nodeInfo.root
		lvl := nodeInfo.level
		sum := nodeInfo.sum
		node.Val = levelSums[lvl] - sum
	}
	return root
}

// Using the pre-defined TreeNode struct
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// depthSum is now a package-level variable
var depthSum []int

// dfs1 performs the first DFS to calculate the sum of node values at each depth
func dfs1(root *TreeNode, d int) {
	if root == nil {
		return // Base case: if the node is null, return
	}

	// If the current depth is greater than or equal to the slice length,
	// add a new element with the current node's value
	if d >= len(depthSum) {
		depthSum = append(depthSum, root.Val)
	} else {
		// Otherwise, add the current node's value to the existing sum at this depth
		depthSum[d] += root.Val
	}

	// Recursively call dfs1 for left and right children, incrementing the depth
	dfs1(root.Left, d+1)
	dfs1(root.Right, d+1)
}

// dfs2 performs the second DFS to replace node values
func dfs2(root *TreeNode, val int, d int) {
	if root == nil {
		return // Base case: if the node is null, return
	}

	// Replace the current node's value with the passed 'val'
	root.Val = val

	// Calculate the sum of cousin nodes' values
	// If there's a next depth, get its sum, otherwise use 0
	c := 0
	if d+1 < len(depthSum) {
		c = depthSum[d+1]
	}
	// Subtract the values of the current node's children (if they exist)
	if root.Left != nil {
		c -= root.Left.Val
	}
	if root.Right != nil {
		c -= root.Right.Val
	}

	// Recursively call dfs2 for left and right children
	// Pass the calculated cousin sum 'c' and increment the depth
	if root.Left != nil {
		dfs2(root.Left, c, d+1)
	}
	if root.Right != nil {
		dfs2(root.Right, c, d+1)
	}
}

// replaceValueInTree is the main function to replace values in the tree
func replaceValueInTree2(root *TreeNode) *TreeNode {
	depthSum = []int{} // Initialize depthSum slice
	// First DFS to calculate depth sums
	dfs1(root, 0)
	// Second DFS to replace values, starting with 0 for the root
	dfs2(root, 0, 0)
	// Return the modified root
	return root
}
