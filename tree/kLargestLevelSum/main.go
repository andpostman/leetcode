package main

import (
	"container/list"
	"fmt"
	"sort"
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
	fmt.Println(kthLargestLevelSum(root, 2))
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var walkTree func(node *TreeNode, level int)
	levelSum := make(map[int]int)
	walkTree = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		levelSum[level] += node.Val
		walkTree(node.Left, level+1)
		walkTree(node.Right, level+1)
	}
	walkTree(root, 0)
	sums := make([]int, 0, len(levelSum))
	for _, val := range levelSum {
		sums = append(sums, val)
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	if k > len(sums) {
		return -1
	} else {
		return int64(sums[k-1])
	}
}

// ANOTHER SOLUTION
func kthLargestLevelSum2(root *TreeNode, k int) int64 {
	// Slice to store the sum of each level
	var levelSums []int64

	// If the tree is empty, return -1
	if root == nil {
		return -1
	}

	// Queue for level-order traversal
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		// Variable to store the sum of the current level
		var levelSum int64

		// Number of nodes at the current level
		pending := queue.Len()
		for i := 0; i < pending; i++ {
			// Process each node at the current level
			node := queue.Remove(queue.Front()).(*TreeNode)

			// Add the node's value to the level sum
			levelSum += int64(node.Val)

			// Add child nodes to the queue for the next level
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		// Add the sum of the current level to the slice
		levelSums = append(levelSums, levelSum)
	}

	// If there are fewer levels than k, return -1
	if len(levelSums) < k {
		return -1
	}

	// Sort the levelSums slice in descending order
	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})

	// Return the kth largest level sum
	return levelSums[k-1]
}
