package main

import "fmt"

func main() {

	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	res := modifiedList([]int{2, 2}, root)
	fmt.Println(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	var iterate func(head *ListNode) *ListNode
	checked := make(map[int]bool)
	for _, num := range nums {
		checked[num] = true
	}
	iterate = func(head *ListNode) *ListNode {
		if head != nil {
			head.Next = iterate(head.Next)
		}
		if head == nil {
			return nil
		}
		if checked[head.Val] {
			return head.Next
		}
		return head
	}
	return iterate(head)
}
