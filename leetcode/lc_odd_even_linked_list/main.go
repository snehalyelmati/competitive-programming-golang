package main

import "fmt"

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	temp := head.Next
	evenHead := head.Next
	original := head

	for head.Next != nil && head.Next.Next != nil {
		head.Next = head.Next.Next
		head = head.Next

		if temp != nil && temp.Next != nil {
			temp.Next = temp.Next.Next
			temp = temp.Next
		}
	}

	head.Next = evenHead
	return original
}

func main() {
	fmt.Println("Hello, world!")
}
