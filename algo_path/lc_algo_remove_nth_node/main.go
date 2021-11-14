package main

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	buf := head
	initial := head
	prev := head

	// 0 1 . 
	// p h n=2

	for i:=0; i<n; i++ {
		buf = buf.Next
	}

	if buf == nil {
		return initial.Next
	}

	// buf.Next != nil
	for buf != nil {
		buf = buf.Next
		prev = head
		head = head.Next
	}
	prev.Next = head.Next

	return initial
}

func main() {
}
