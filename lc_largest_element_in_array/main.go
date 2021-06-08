package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	nums = mergeSort(nums)
	return nums[len(nums)-k]
}

func mergeSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	mid := len(n)/2
	left := mergeSort(n[:mid])
	right := mergeSort(n[mid:])

	return merge(left, right)
}

func merge(l []int, r []int) []int {

	res := make([]int, len(l) + len(r))
	li, ri, k := 0, 0, 0
	for li<len(l) && ri<len(r) {
		if l[li] < r[ri] {
			res[k] = l[li]
			li++
		} else {
			res[k] = r[ri]
			ri++
		}
		k++
	}

	for li < len(l) {
		res[k] = l[li]
		k++
		li++
	}

	for ri < len(r) {
		res[k] = r[ri]
		k++
		ri++
	}

	return res
}

func main() {
	// nums := []int{3,2,1,5,6,4}
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}