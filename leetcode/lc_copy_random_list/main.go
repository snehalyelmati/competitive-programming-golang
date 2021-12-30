package main

import (
	"fmt"
	"reflect"
)

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	initial := head
	var dummyNode Node
	prev := &dummyNode
	newHead := &dummyNode

	hMap := make(map[*Node]*Node)
	for head != nil {
		// create a new node
		// assign value
		newNode := Node{Val: head.Val, Next: nil, Random: nil}

		// point prev.Next to this node
		prev.Next = &newNode

		// point it prev
		prev = &newNode

		// save the nodes mapping in a hash map
		hMap[head] = &newNode

		// move head to the next node
		head = head.Next
	}
	// point prev to nil
	prev.Next = nil

	// pointing it to the head of the new list
	newHead = newHead.Next

	// iterate through the old list and point all the random pointers
	head = initial
	newInitial := newHead
	for head != nil {
		newHead.Random, _ = hMap[head.Random]
		head = head.Next
		newHead = newHead.Next
	}

	return newInitial
}

func main() {
	head := &Node{Val: 3, Next: nil, Random: nil}
	fmt.Println(reflect.TypeOf(head))
	fmt.Println(copyRandomList(head))
}
