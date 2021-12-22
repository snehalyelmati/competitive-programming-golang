package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil && head.Next != nil {
		first := head.Next
		second := head.Next.Next
		first.Next = head
		head.Next = prev
		head = second
		prev = first
	}

	if head != nil {
		head.Next = prev
		prev = head
	}

	return prev
}

func main() {
}
