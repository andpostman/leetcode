package splitLinkedListInParts

import "fmt"

func main() {

	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  10,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	res := splitListToParts(root, 3)
	fmt.Println(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	resultArr := make([]*ListNode, k)
	if head == nil {
		return resultArr
	}
	countOfNodes := 0
	for h := head; h != nil; h = h.Next {
		countOfNodes++
	}
	parts := countOfNodes / k
	remainigParts := countOfNodes % k
	var splitList func(h *ListNode)
	iter := 0
	splitList = func(h *ListNode) {
		length := parts
		if remainigParts > 0 {
			length++
			remainigParts--
		}
		var iterate func(*ListNode, int) (current *ListNode, next *ListNode)
		iterate = func(hh *ListNode, l int) (current *ListNode, next *ListNode) {
			if l > 0 && hh != nil {
				l--
				hh.Next, next = iterate(hh.Next, l)
			} else if hh == nil {
				return
			} else {
				next = hh
				return
			}
			current = hh
			return
		}
		res, next := iterate(h, length)
		resultArr[iter] = res
		if next != nil {
			iter++
			splitList(next)
		}
	}
	splitList(head)
	return resultArr
}
