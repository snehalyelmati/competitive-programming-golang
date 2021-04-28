package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//   	3
//    /	  \
//   9  	20
//   	  /	   \
// 	     15		7
//
// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0, 0, len(preorder)-1, preorder, inorder)
}

func helper(preStart, inStart, inEnd int, preorder, inorder []int) *TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}

	var head TreeNode
	head.Val = preorder[preStart]

	id_root := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == head.Val {
			id_root = i
			break
		}
	}

	head.Left = helper(preStart+1, inStart, id_root-1, preorder, inorder)
	head.Right = helper(preStart+id_root-inStart+1, id_root+1, inEnd, preorder, inorder)

	return &head
}

func main() {
	fmt.Println("Hello, world!")
}
