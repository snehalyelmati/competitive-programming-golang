package main

import (
	"fmt"
)

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	currList := make([]*TreeNode, 0)
	res := make([][]int, 0)
	flag := false

	currList = append(currList, root)
	for len(currList) != 0 {
		fmt.Println(currList, res, flag)

		size := len(currList)
		temp := make([]int, 0)
		for i := 0; i < size; i++ {
			temp = append(temp, currList[0].Val)
			if currList[0].Left != nil {
				currList = append(currList, currList[0].Left)
			}
			if currList[0].Right != nil {
				currList = append(currList, currList[0].Right)
			}
			currList = currList[1:]
		}
		if flag {
			temp2 := make([]int, 0)
			for i := len(temp) - 1; i >= 0; i-- {
				temp2 = append(temp2, temp[i])
			}
			copy(temp, temp2)
			fmt.Println(temp)
		}
		res = append(res, temp)
		flag = !flag
	}

	return res
}

func main() {
	fmt.Println("Hello, world!")
}
