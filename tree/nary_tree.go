package main

import "fmt"

func main() {

	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val:      5,
						Children: nil,
					},
					{
						Val:      6,
						Children: nil,
					},
				},
			},
			{
				Val:      2,
				Children: nil,
			},
			{
				Val:      4,
				Children: nil,
			},
		},
	}
	res := postorder(root)
	fmt.Println(res)

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return nil
	}
	if root.Children != nil {
		for _, child := range root.Children {
			res = append(res, postorder(child)...)
		}
	}
	res = append(res, root.Val)
	return res
}
