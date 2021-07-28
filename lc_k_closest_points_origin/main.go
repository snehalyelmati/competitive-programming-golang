package main

import (
	"fmt"
	"sort"
)

// using sort interface
type EuPoints [][]int
func (e EuPoints) Len() int {
	return len(e)
}
func (e EuPoints) Less(i, j int) bool {
	return e[i][0]*e[i][0] + e[i][1]*e[i][1] < e[j][0]*e[j][0] + e[j][1]*e[j][1]
}
func (e EuPoints) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func kClosest(p [][]int, k int) [][]int {
	sort.Sort(EuPoints(p))
	return p[:k]
}

// Heap solution - not complete
func kClosestOld2(p [][]int, k int) [][]int {
	// implement max-heap
	// whenever the size is greater than k, pop the root node
	// insert at the end, then heapify
	maxHeap := make([]int, k)

	// array with distances
	radiusArray := make([]int, len(p))
	for k, v := range p {
		radiusArray[k] = v[0]*v[0] + v[1]*v[1]
	}

	idx := -1
	for _, v := range radiusArray {
		idx++
		if idx >= k {
			// replace the 0th element with the new one
			maxHeap[0] = v
			fmt.Println("Reached limit manipulating maxHeap...")
			fmt.Println("Before: ", maxHeap)

			// heapify(&maxHeap)
			// check if the root node is less than its children
			// if it is swap with the minimum element
			i := 0
			for (i*2+1 < k && maxHeap[i] < maxHeap[i*2+1]) || (i*2+2 < k && maxHeap[i] < maxHeap[i*2+2]) {
				tempIdx := i
				if maxHeap[i] < maxHeap[i*2+1] {
					maxHeap[i], maxHeap[i*2+1] = maxHeap[i*2+1], maxHeap[i]
					tempIdx = i*2+1
				}
				if maxHeap[i] < maxHeap[i*2+2] {
					maxHeap[i], maxHeap[i*2+2] = maxHeap[i*2+2], maxHeap[i]
					tempIdx = i*2+2
				}
				if tempIdx == i {
					break
				}
			}

			fmt.Println("After: ", maxHeap)
		} else {
			fmt.Println("Inserting elements at the right position...")
			// insert the element at the end
			// bubble it up
			maxHeap[idx] = v
			// save the index
			temp := idx
			// parent < child => swap
			for idx < k &&  maxHeap[idx/2-1] < maxHeap[idx] {
				maxHeap[idx/2-1], maxHeap[idx] = maxHeap[idx], maxHeap[idx/2-1]
				idx = idx/2-1
			}
			idx = temp
		}
	}

	return nil
}

func kClosestOld(p [][]int, k int) [][]int {
	radiusArray := make([]int, len(p))
	for k, v := range p {
		radiusArray[k] = v[0]*v[0] + v[1]*v[1]
	}

	// fmt.Println(p)
	// fmt.Println(radiusArray)
	sort.SliceStable(radiusArray, func(i, j int) bool {
		if radiusArray[i] < radiusArray[j] {
			p[i], p[j] = p[j], p[i]
			return true
		}
		return false
	})
	// fmt.Println(p)
	// fmt.Println(radiusArray)

	return p[:k]
}

func main() {
	// points := [][]int {
	// 	{3,2}, {-2,2},
	// 	{2,2}, {2,-2},
	// 	{4,2}, {-2,-2},
	// 	{1,2}, {2,2},
	// }
	points := [][]int {{1,3}, {-2,-2}}
	k := 1

	fmt.Println(kClosest(points, k))

	// {{572,433},{-601,-305},{-76,714},{-143,500},{659,294},{-56,-751},{624,-785},{-402,458},{93,-552},{493,985},{877,-351},{10,706},{-904,-632},{710,982}}
	// {{-422,117},{443,23},{240,454},{-143,500},{-301,-466},{93,-552},{-402,458},{500,-376},{-38,-662},{-658,97},{-37,665},{-601,-305},{10,706},{-106,-701}}
}

