package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var slowHead *ListNode = head
	for head != nil && head.Next != nil {
		slowHead = slowHead.Next
		head = head.Next.Next
	}

	return slowHead
}

func main() {
}
