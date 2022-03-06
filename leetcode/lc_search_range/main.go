package main

import "fmt"

func searchRange(nums []int, target int) []int {
	l, h := 0, len(nums)-1
	if h == -1 {
		return []int{-1, -1}
	}

	idx := binarySearch(nums, l, h, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	l, h = idx, idx
	for l >= 0 && nums[l] == target {
		l--
	}
	l++
	for h < len(nums) && nums[h] == target {
		h++
	}
	h--

	return []int{l, h}
}

func binarySearch(n []int, l, h, target int) int {
	mid := (l + h) / 2

	for l <= h {
		fmt.Println(l, h)
		if n[mid] < target {
			return binarySearch(n, mid+1, h, target)
		} else if n[mid] > target {
			return binarySearch(n, l, mid-1, target)
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	nums := []int{1}
	target := 1
	// fmt.Println(searchRange(nums, target))
	fmt.Println(binarySearch(nums, 0, len(nums)-1, target))
}
