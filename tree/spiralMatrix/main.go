package main

import "fmt"

func main() {
	root := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 2,
											Next: &ListNode{
												Val: 5,
												Next: &ListNode{
													Val: 5,
													Next: &ListNode{
														Val:  0,
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
				},
			},
		},
	}
	res := spiralMatrix(3, 5, root)
	for i := range res {
		for j := range res[i] {
			fmt.Printf("[%d] ", res[i][j])
		}
		fmt.Println()
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	vectors := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// processed := make(map[string]struct{})
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}
	idx := 0
	current := head
	for i, j := 0, 0; current != nil; {
		matrix[i][j] = current.Val
		current = current.Next

		nextI := i + vectors[idx][0]
		nextJ := j + vectors[idx][1]

		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || matrix[nextI][nextJ] != -1 {
			idx = (idx + 1) % 4
			nextI = i + vectors[idx][0]
			nextJ = j + vectors[idx][1]
		}
		i, j = nextI, nextJ
	}
	return matrix
}
