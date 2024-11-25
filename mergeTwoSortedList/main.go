package main

import "fmt"

func main() {
	fmt.Println(mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	res := &ListNode{}
	if list1.Val < list2.Val {
		res.Val = list1.Val
		res.Next = mergeTwoLists(list1.Next, list2)
	} else {
		res.Val = list2.Val
		res.Next = mergeTwoLists(list1, list2.Next)
	}
	return res
}
