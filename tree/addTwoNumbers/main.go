package main

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	addTwoNumbers(l1, l2)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numbersL1 := appendAll(l1)
	numbersL2 := appendAll(l2)
	res := makeResult(numbersL1, numbersL2)
	result := reverseConvertToNode(res)
	return result
}

func appendAll(l *ListNode) []int {
	var values []int
	var getNumbers func(next *ListNode)
	getNumbers = func(next *ListNode) {
		values = append(values, next.Val)
		if next.Next != nil {
			getNumbers(next.Next)
		}
	}
	getNumbers(l)
	return values
}

func reverseConvertToNode(numbers []int) *ListNode {
	var appendNodes func(int) *ListNode
	appendNodes = func(i int) *ListNode {
		var node *ListNode
		if i < len(numbers) {
			node = new(ListNode)
			node.Val = numbers[i]
			i++
			node.Next = appendNodes(i)
		}
		return node
	}
	node := appendNodes(0)
	return node
}

func makeResult(l1, l2 []int) []int {
	var resNum []int
	switch lens := len(l1) - len(l2); {
	case lens == 0:
		resNum = appendResNum(true, l1, l2)
	case lens > 0:
		resNum = appendResNum(false, l1, l2)
	default:
		resNum = appendResNum(false, l2, l1)
	}
	return resNum
}

func appendResNum(isEq bool, moreArr, lessArr []int) []int {
	var resNum []int
	needToAdd := false
	if isEq {
		for i := 0; i < len(moreArr); i++ {
			numL1 := moreArr[i]
			numL2 := lessArr[i]
			if needToAdd {
				numL1++
				needToAdd = false
			}
			sum := numL1 + numL2
			if sum/10 > 0 {
				needToAdd = true
			}
			resNum = append(resNum, sum%10)
		}
	} else {
		for i := 0; i < len(lessArr); i++ {
			numL1 := moreArr[i]
			numL2 := lessArr[i]
			if needToAdd {
				numL1++
				needToAdd = false
			}
			sum := numL1 + numL2
			if sum/10 > 0 {
				needToAdd = true
			}
			resNum = append(resNum, sum%10)
		}
		for i := len(lessArr); i < len(moreArr); i++ {
			numL2 := moreArr[i]
			if needToAdd {
				numL2++
				needToAdd = false
			}
			sum := numL2
			if sum/10 > 0 {
				needToAdd = true
			}
			resNum = append(resNum, sum%10)
		}
	}
	if needToAdd {
		resNum = append(resNum, 1)
	}
	return resNum
}
