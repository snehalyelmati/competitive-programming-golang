package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	nodes := make([]TreeNode, 0)
	nodes = append(nodes, *root)

	for len(nodes) != 0 {
		size := len(nodes)

		arr := make([]int, 0)
		for i := 0; i < size; i++ {
			arr = append(arr, nodes[i].Val)

			if nodes[i].Left != nil {
				nodes = append(nodes, *nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, *nodes[i].Right)
			}
		}
		res = append(res, arr)

		nodes = nodes[size:]
	}

	return res
}

func main() {
	fmt.Println("Hello, world!")
}
