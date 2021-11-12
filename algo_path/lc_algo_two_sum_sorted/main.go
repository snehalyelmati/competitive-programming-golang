package main

import "fmt"

func twoSum(nums []int, target int) []int {
	f := 0
	l := len(nums) - 1

	for f != l {
		if nums[f]+nums[l] == target {
			return []int{f + 1, l + 1}
		} else if nums[f]+nums[l] > target {
			l--
		} else if nums[f]+nums[l] < target {
			f++
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
