package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	fmt.Println(nums)
	fmt.Println(l, h)

	return l
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0

	fmt.Println(searchInsert(nums, target))
}
