package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mainLogic(root *TreeNode) []int {
	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root != nil {
		return mainLogic(root)
	} 
	return nil
}

func main() {
	fmt.Println("Hello, world!")
}
