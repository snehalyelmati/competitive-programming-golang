package main

import "fmt"

func searchOld(nums []int, target int) int {
	l := 0
	h := len(nums) - 1
	m := int((l + h) / 2)

	if target == nums[l] {
		return l
	} else if target == nums[h] {
		return h
	}

	for i := l; i <= h; i++ {
		if target == nums[m] {
			return m
		}
		if target < nums[m] {
			h = m
		} else {
			l = m
		}
		m = int((l + h) / 2)
	}
	return -1
}

func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := (i + j) / 2
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			i += 1
		} else {
			j -= 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Printf("Array: %v, Target: %v, Index: %v\n", nums, target, search(nums, target))
}
