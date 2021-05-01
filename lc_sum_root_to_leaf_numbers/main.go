package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum := 0
	if root.Left != nil {
		sum += helper(root.Left, root.Val)
	}
	if root.Right != nil {
		sum += helper(root.Right, root.Val)
	}
	return sum
}

func helper(root *TreeNode, prior int) int {
	if root.Left == nil && root.Right == nil {
		return prior*10 + root.Val
	}

	sum := 0
	if root.Left != nil {
		sum += helper(root.Left, prior*10+root.Val)
	}

	if root.Right != nil {
		sum += helper(root.Right, prior*10+root.Val)
	}

	return sum
}

func main() {
	fmt.Println("Hello, world!")
}
