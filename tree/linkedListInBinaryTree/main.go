package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSubPath(head, root))
}

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var walkTree func(h *ListNode, r *TreeNode) bool
	walkTree = func(listNode *ListNode, treeNode *TreeNode) bool {
		if listNode == nil {
			return true
		}
		if treeNode == nil {
			return false
		}
		isPath := false
		if listNode.Val == treeNode.Val {
			isPath = walkTree(listNode.Next, treeNode.Right) || walkTree(listNode.Next, treeNode.Left)
		}
		return isPath
	}
	if root == nil {
		return false
	}

	isPath := walkTree(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
	return isPath
}
