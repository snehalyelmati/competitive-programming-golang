package main

import (
	"fmt"
)

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	nums = merge(mergeSort(nums[:len(nums)/2]), mergeSort(nums[len(nums)/2:]))
	return nums
}

func merge(a, b []int) []int {
	fmt.Println(a, b)
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, b[j])
		j++
	}

	return res
}

func main() {
	// nums := []int{8, 3, 4, 5, 12, 6}
	nums := []int{}
	fmt.Println(mergeSort(nums))
}
