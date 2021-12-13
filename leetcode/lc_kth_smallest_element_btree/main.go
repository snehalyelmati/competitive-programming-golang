package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var counter int

func inorderArr(root *TreeNode) []int {
	arr := make([]int, 0)

	if root == nil {
		return nil
	}

	arr = append(arr, inorderArr(root.Left)...)

	arr = append(arr, root.Val)
	counter--
	if counter == 0 {
		return arr
	}

	arr = append(arr, inorderArr(root.Right)...)

	return arr
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)

	counter = k
	arr = append(arr, inorderArr(root)...)

	return arr[k-1]
}

func main() {
	fmt.Println("Hello, world!")
}
