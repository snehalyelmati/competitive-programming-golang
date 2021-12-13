package main

import (
	"fmt"
	"math"
)

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hMap := make(map[*ListNode]bool)

	for headA != nil {
		hMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	hMap := make(map[*ListNode]bool)
	for headA != nil && headB != nil {
		if _, ok := hMap[headA]; ok {
			return headA
		}
		hMap[headA] = true
		headA = headA.Next

		if _, ok := hMap[headB]; ok {
			return headB
		}
		hMap[headB] = true
		headB = headB.Next
	}

	for headA != nil {
		if _, ok := hMap[headA]; ok {
			return headA
		}
		hMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hMap[headB]; ok {
			return headB
		}
		hMap[headB] = true
		headB = headB.Next
	}

	fmt.Println(hMap)

	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	countA := 0
	countB := 0

	tempA := headA
	tempB := headB

	for tempA != nil {
		countA++
		tempA = tempA.Next
	}

	for tempB != nil {
		countB++
		tempB = tempB.Next
	}

	diff := int(math.Abs(float64(countA - countB)))

	if countA >= countB {
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}

		// check
		for headA != nil {
			if headA == headB {
				return headA
			}
			if headA.Next == headB.Next {
				return headA.Next
			}
			headA = headA.Next
			headB = headB.Next
		}

		return nil
	} else if countB > countA {
		for i := 0; i < diff; i++ {
			headB = headB.Next
		}

		// check
		for headB != nil {
			if headA == headB {
				return headA
			}
			if headA.Next == headB.Next {
				return headA.Next
			}
			headA = headA.Next
			headB = headB.Next
		}

		return nil
	}

	return nil
}

func getIntersectionNode4(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}

	return a
}

func main() {
	fmt.Println("Hello, world!")
}
