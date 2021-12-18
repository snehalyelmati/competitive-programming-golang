package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	// fmt.Println(nums)
	i, j := 0, 0
	size := len(nums)
	criteria := len(nums) / 3

	res := make([]int, 0)
	count := 0
	for i < size && j < size {
		// fmt.Println(nums[i], nums[j], i, j)
		if nums[i] == nums[j] {
			j++
			count++
		} else {
			// fmt.Println(count, criteria)
			if count > criteria {
				res = append(res, nums[i])
			}
			i = j
			count = 0
		}
	}

	if count > criteria {
		res = append(res, nums[i])
	}

	return res
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
