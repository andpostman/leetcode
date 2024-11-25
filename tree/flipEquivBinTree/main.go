package main

import "fmt"

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(flipEquiv(root1, root2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(r1, r2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		return node1.Val == node2.Val &&
			((dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)) ||
				(dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)))
	}
	return dfs(root1, root2)
}
