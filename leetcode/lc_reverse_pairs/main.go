package main

import (
	"fmt"
)

func mergeSort(nums []int, count int) ([]int, int) {
	if len(nums) == 1 {
		return nums, count
	}

	mid := len(nums) / 2
	left, count := mergeSort(nums[:mid], count)
	right, count := mergeSort(nums[mid:], count)

	return merge(left, right, count)
}

func merge(left, right []int, count int) ([]int, int) {
	i := 0
	j := 0

	for i = 0; i < len(left); i++ {
		for j < len(right) && left[i] > 2*right[j] {
			j++
		}
		count += j
	}

	i, j = 0, 0
	res := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	if i != len(left) {
		res = append(res, left[i:]...)
	}
	if j != len(right) {
		res = append(res, right[j:]...)
	}

	return res, count
}

func reversePairs(nums []int) int {
	sortedArray, count := mergeSort(nums, 0)
	fmt.Println(sortedArray)
	return count
}

func main() {
	nums := []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs(nums))
}
