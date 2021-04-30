package main

import "fmt"

//   Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectOld(root *Node) *Node {
	if root == nil {
		return nil
	}

	var que []*Node
	que = append(que, root)

	for len(que) != 0 {
		var temp []*Node
		size := len(que)
		for i := 0; i < size; i++ {
			if que[0].Left != nil {
				temp = append(temp, que[0].Left)
				que = append(que, que[0].Left)
			}
			if que[0].Right != nil {
				temp = append(temp, que[0].Right)
				que = append(que, que[0].Right)
			}
			que = que[1:]
		}

		for i := 0; i < len(temp)-1; i++ {
			temp[i].Next = temp[i+1]
		}
	}

	return root
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		root.Left.Next = root.Right
		connect(root.Left)
	}

	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect(root.Right)
	}

	return root
}

func main() {
	fmt.Println("Hello, world!")
}
