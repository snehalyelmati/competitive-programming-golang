package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hMap := make(map[int]bool)
	for _, d := range nums {
		hMap[d] = true
	}

	maxCount := 1
	for i := 0; i < len(nums); i++ {
		if _, ok := hMap[nums[i]-1]; ok {
			continue
		}
		temp := nums[i] + 1
		count := 1
		for _, ok := hMap[temp]; ok; {
			count++
			if maxCount < count {
				maxCount = count
			}
			temp++
			_, ok = hMap[temp]
		}
	}

	return maxCount
}

func longestConsecutiveOld(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	maxCount := 1
	sort.Ints(nums)
	fmt.Println(nums)

	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else if nums[i] < nums[i+1] {
			count = 1
		}
	}

	return maxCount
}

func main() {
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
