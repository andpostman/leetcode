package main

import "fmt"

func main() {
	root := &ListNode{
		Val: 18,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	res := insertGreatestCommonDivisors(root)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var iterateNode func(node *ListNode) *ListNode
	iterateNode = func(node *ListNode) *ListNode {
		if node != nil {
			node.Next = iterateNode(node.Next)
		}
		if node == nil {
			return nil
		}
		if node.Next == nil {
			return node
		} else {
			divisors := gcd(node.Val, node.Next.Val)
			commonDivisorsNode := &ListNode{
				Val:  divisors,
				Next: node.Next,
			}
			node.Next = commonDivisorsNode
			return node
		}
	}
	return iterateNode(head)
}

// gcd calculates the greatest common divisor (GCD) of two integers using the Euclidean algorithm.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// func nod(a, b int) int {
// 	if a%b == 0 {
// 		return b
// 	} else if b%a == 0 {
// 		return a
// 	}
// 	lower := 0
// 	if a < b {
// 		lower = a
// 	} else {
// 		lower = b
// 	}
// 	maxNod := 0
// 	for i := 1; i <= lower; i++ {
// 		if a%i == 0 && b%i == 0 {
// 			maxNod = i
// 		}
// 	}
// 	return maxNod
// }
