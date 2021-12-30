package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	initial := head
	temp := head

	length := 0
	for temp != nil {
		temp = temp.Next
		length++
	}
	k %= length
	temp = head

	for k > 0 {
		temp = temp.Next
		k--
	}

	for temp.Next != nil {
		head = head.Next
		temp = temp.Next
	}
	fmt.Println(initial.Val, head.Val, temp.Val)

	if head.Next != nil {
		tail := head
		head = head.Next
		temp.Next = initial
		tail.Next = nil
	} else {
		head = initial
	}

	return head
}

func main() {
}

//	1	2	3	4	5
//	-	h	-	-	-
//	-	t	-	-	-
