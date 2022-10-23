package main

import "fmt"

func findPeakElement(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		mid := (l + h) / 2
		fmt.Println(mid)
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return l
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Peak element in the array:", nums, findPeakElement(nums))
}
