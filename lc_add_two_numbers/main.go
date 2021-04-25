package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	buff := 0
	var head *ListNode
	var buffNode ListNode

	count1 := 0
	count2 := 0
	templ1 := l1
	templ2 := l2

	for templ1 != nil {
		count1++
		templ1 = templ1.Next
	}

	for templ2 != nil {
		count2++
		templ2 = templ2.Next
	}

	if count2 > count1 {
		l1, l2 = l2, l1
	}

	if l1 != nil {
		head = l1
	} else if l2 != nil {
		head = l2
	} else {
		return nil
	}
	temp := head

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + buff
		fmt.Println(l1.Val, l2.Val, sum, buff)
		head.Val = (sum) % 10
		buff = (sum) / 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			head.Next = nil
			break
		}

		if l1 != nil {
			head = l1
		} else {
			head = l2
		}
	}

	for l1 != nil {
		sum := l1.Val + buff
		fmt.Println(l1.Val, sum, buff, "l1")
		head.Val = sum % 10
		buff = sum / 10
		l1 = l1.Next
		if l1 != nil {
			head = l1
		}
	}

	// for l2 != nil {
	// 	sum := l2.Val + buff
	// 	fmt.Println(l2.Val, sum, buff, "l2")
	// 	head.Val = sum % 10
	// 	buff = sum / 10
	// 	l2 = l2.Next
	// 	if l2 != nil {
	// 		head = l2
	// 	}
	// }

	if buff != 0 {
		buffNode.Val = buff
		buffNode.Next = nil
		head.Next = &buffNode
	}

	// temp2 := temp
	// for temp != nil {
	// 	fmt.Println(temp)
	// 	temp = temp.Next
	// }

	return temp
}

func main() {
	fmt.Println("Hello, world!")
}
