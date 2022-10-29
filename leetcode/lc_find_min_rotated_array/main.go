package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	// where it starts decreasing
	pivot := findPivot(nums)
	return nums[(pivot+1)%len(nums)]
}

func findPivot(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h && nums[l] > nums[h] {
		mid := l + (h-l)/2
		if nums[mid] > nums[l] {
			l = mid
		} else if nums[mid] <= nums[l] {
			h = mid
		}
	}
	return h
}

func binarySearch(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2, 3}
	fmt.Println(findMin(nums))
}
