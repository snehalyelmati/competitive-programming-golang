package main

import "fmt"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)

	if size == 0 {
		return nil
	} else if size == 1 {
		return &TreeNode{nums[0], nil, nil}
	} else if size == 2 {
		leftNode := TreeNode{nums[0], nil, nil}
		return &TreeNode{nums[1], &leftNode, nil}
	} else if size == 3 {
		leftNode := TreeNode{nums[0], nil, nil}
		rightNode := TreeNode{nums[2], nil, nil}
		return &TreeNode{nums[1], &leftNode, &rightNode}
	}

	leftNode := sortedArrayToBST(nums[:size/2])
	rightNode := sortedArrayToBST(nums[size/2+1:])
	rootNode := TreeNode{nums[size/2], leftNode, rightNode}
	return &rootNode
}

func main() {
	fmt.Println("Hello, world!")
}
