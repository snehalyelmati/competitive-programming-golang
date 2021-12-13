package main

import "fmt"

func searchInsert(nums []int, target int) int {
	size := len(nums)
	if target > nums[size-1] {
		return size
	} else if target <= nums[0] {
		return 0
	} else if target == nums[size-1] {
		return size - 1
	}

	m := 0
	for l, h := 0, size-1; l < h; {
		m = (l + h) / 2
		if nums[m] == target {
			return m
		}

		if target > nums[m] {
			l += 1
		} else {
			h -= 1
		}
	}
	return m + 1
}

func main() {
	nums1 := []int{1, 3, 6, 7}
	target1 := 5

	nums2 := []int{1, 3, 6, 7}
	target2 := 2

	nums3 := []int{1, 3, 6, 7}
	target3 := 7

	nums4 := []int{1, 3, 6, 7}
	target4 := 0

	nums5 := []int{1}
	target5 := 0

	fmt.Printf("Nums: %v, target: %v, id: %v\n", nums1, target1, searchInsert(nums1, target1))
	fmt.Printf("Nums: %v, target: %v, id: %v\n", nums2, target2, searchInsert(nums2, target2))
	fmt.Printf("Nums: %v, target: %v, id: %v\n", nums3, target3, searchInsert(nums3, target3))
	fmt.Printf("Nums: %v, target: %v, id: %v\n", nums4, target4, searchInsert(nums4, target4))
	fmt.Printf("Nums: %v, target: %v, id: %v\n", nums5, target5, searchInsert(nums5, target5))
}
