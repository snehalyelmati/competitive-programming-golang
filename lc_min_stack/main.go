package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	intStack []int
	min      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack MinStack
	stack.intStack = make([]int, 0)
	stack.min = math.MaxInt32
	return stack
}

func (this *MinStack) Push(val int) {
	if this.min > val {
		this.min = val
	}
	this.intStack = append(this.intStack, val)
}

func (this *MinStack) Pop() {
	if this.min == this.Top() {
		this.intStack = this.intStack[:len(this.intStack)-1]
		this.min = findMin(this.intStack)
	} else {
		this.intStack = this.intStack[:len(this.intStack)-1]
	}
}

func findMin(a []int) int {
	min := math.MaxInt32
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}

func (this *MinStack) Top() int {
	return this.intStack[len(this.intStack)-1]
}

func (this *MinStack) GetMin() int {
	// minEle := this.intStack[0]
	// for _, v := range this.intStack {
	// 	if v < minEle {
	// 		minEle = v
	// 	}
	// }
	// return minEle
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	var minStack MinStack = Constructor()
	fmt.Println(minStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
