package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil { 
		return false
	}

	if isSame(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil { 
		return false
	}

	return root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}

func isSubtreeOld(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil && subRoot != nil || root != nil && subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val && isSubtree(root.Left, subRoot.Left) && isSubtree(root.Right, subRoot.Right) {
		return true
	}

	if root.Left != nil && isSubtree(root.Left, subRoot) {
		return true
	}
	if root.Right != nil && isSubtree(root.Right, subRoot) {
		return true
	}

	return false
}

func main() {
	fmt.Println("Hello, world!")
}