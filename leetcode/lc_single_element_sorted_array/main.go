package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	l, h := 0, len(nums)-2
	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] == nums[mid^1] { // left
			l = mid + 1
		} else { // right
			h = mid - 1
		}
	}
	return nums[l]
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 8, 8, 9, 9}
	fmt.Println(singleNonDuplicate(nums))
}

// 1 1 2 2 3 3|4 8 8
// 0 1 0 1 0 1 0 1 0
