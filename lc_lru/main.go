package main

import (
	"fmt"
)

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

type LRUCache struct {
	kv       map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := Node{
		prev:  nil,
		next:  nil,
		key:   0,
		value: 0,
	}

	tail := Node{
		prev:  nil,
		next:  nil,
		key:   0,
		value: 0,
	}

	head.next = &tail
	tail.prev = &head

	kv := make(map[int]*Node, capacity)

	return LRUCache{
		head:     &head,
		tail:     &tail,
		kv:       kv,
		capacity: capacity,
	}
}

func printLinkedList(h *Node) {
	if h != nil {
		fmt.Print("(", h.key, ",", h.value, ")")
		h = h.next
	}
	for h != nil {
		fmt.Print(" -> ", "(", h.key, ",", h.value, ")")
		h = h.next
	}
	fmt.Println()
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("Get: ", key)
	if node, ok := this.kv[key]; ok {
		this.remove(node)
		this.insert(node)
		printLinkedList(this.head)
		// fmt.Printf("GET.kv: %v\n", this.kv)
		return node.value
	}
	printLinkedList(this.head)
	// fmt.Printf("GET.kv: %v\n", this.kv)
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("Put: ", key, value)
	if node, ok := this.kv[key]; ok {
		this.remove(node)
		node.value = value
	}
	this.insert(&Node{
		prev:  nil,
		next:  nil,
		key:   key,
		value: value,
	})
	// fmt.Printf("PUT.kv: %v\n", this.kv)
	printLinkedList(this.head)
}

// remove from bottom
func (this *LRUCache) remove(node *Node) {
	fmt.Printf("remove node: %v\n", node)
	fmt.Printf("this.capacity: %v\n", this.capacity)
	printLinkedList(node)
	fmt.Println(node.prev)
	fmt.Println(node.prev.next)
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(this.kv, node.key)
	this.capacity++
	// fmt.Printf("this.kv: %v\n", this.kv)
}

// insert at top
func (this *LRUCache) insert(node *Node) {
	fmt.Printf("insert node: %v\n", node)
	fmt.Printf("this.capacity: %v\n", this.capacity)
	if this.capacity == 0 {
		this.remove(this.tail.prev)
	}
	if this.head.next == this.tail {
		this.head.next = node
		node.prev = this.head

		node.next = this.tail
		this.tail.prev = node
	} else {
		node.next = this.head.next
		node.next.prev = node
		this.head.next = node
		node.prev = this.head
	}
	this.kv[node.key] = node
	this.capacity--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	fmt.Println("Hello, world!")
	capacity := 2
	obj := Constructor(capacity)
	fmt.Println()
	fmt.Println(obj.Get(2))
	fmt.Println()
	obj.Put(2, 6)
	fmt.Println()
	fmt.Println(obj.Get(1))
	fmt.Println()
	obj.Put(1, 5)
	fmt.Println()
	obj.Put(1, 2)
	fmt.Println()
	fmt.Println(obj.Get(1))
	fmt.Println()
	fmt.Println(obj.Get(2))
}
