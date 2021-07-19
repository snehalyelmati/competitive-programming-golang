package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	eles map[int]int
	list []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	hmap := make(map[int]int)
	newList := make([]int, 0)

	return RandomizedSet{
		eles: hmap,
		list: newList,
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.eles[val]; !ok {
		this.eles[val] = len(this.list)
		this.list = append(this.list, val)
		fmt.Println("In insert()", val, ":", this.list, this.eles)
		return true
	}

	fmt.Println("In insert()", val, ":", this.list, this.eles)
	return false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	listSize := len(this.list)
	if idx, ok := this.eles[val]; ok {
		this.list[listSize-1], this.list[idx] = this.list[idx], this.list[listSize-1]

		// fmt.Println(this.list, this.list[listSize-1], this.list[idx], idx)
		// fmt.Println("Before eles:", this.eles)
		// _, found := this.eles[this.list[listSize-1]]
		// fmt.Println("removing:", this.list[listSize-1], found)

		delete(this.eles, val)

		this.list = this.list[:listSize-1]

		fmt.Println("In ok remove", val, ":", this.list, this.eles)
		return true
	}
	fmt.Println("In remove", val, ":", this.list, this.eles)

	return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	fmt.Println("In random():", this.list)
	// mapkeys := reflect.ValueOf(this.eles).MapKeys()
	// return int(mapkeys[rand.Intn(len(mapkeys))].Int())

	randIdx := rand.Intn(len(this.list))
	return this.list[randIdx]
}

/**
* Your RandomizedSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
*/

func main(){
	fmt.Println("Hello, world!")

	// Your RandomizedSet object will be instantiated and called as such:
	obj := Constructor();
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())

	fmt.Println()

	obj1 := Constructor();
	fmt.Println(obj1.Insert(1))
	fmt.Println(obj1.Remove(2))
	fmt.Println(obj1.Insert(2))
	fmt.Println(obj1.GetRandom())
	fmt.Println(obj1.Remove(1))
	fmt.Println(obj1.Insert(2))
	fmt.Println(obj1.GetRandom())

	obj2 := Constructor();
	fmt.Println(obj2.Remove(0))
	fmt.Println(obj2.Remove(0))
	fmt.Println(obj2.Insert(0))
	fmt.Println(obj2.GetRandom())
	fmt.Println(obj2.Remove(0))
	fmt.Println(obj2.Insert(0))
}
