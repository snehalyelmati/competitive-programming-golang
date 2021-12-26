package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	initial := head
	// find the length of the linked list
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	prev := head // prev is nil
	head = initial
	target := head

	for i := 0; i+k <= length; i += k {
		// flip a group of k elements
		count := 0

		// find the last element in the group
		// reverse all the arrows
		curprev := prev
		for count < k {
			first := target.Next
			target.Next = curprev
			curprev = target
			target = first

			count++
		}

		// plug it in the linked list
		head.Next = target
		if prev != nil {
			prev.Next = curprev
		} else {
			initial = curprev
		}
		prev = head
		head = target
	}

	return initial
}

func main() {
}
