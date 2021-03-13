package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func singleNumberOld2(nums []int) int {
	size := len(nums)

	hMap := make(map[int]bool, size)
	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			hMap[v] = false
		} else {
			hMap[v] = true
		}
	}
	for _, v := range nums {
		if data, _ := hMap[v]; data {
			return v
		}
	}
	return nums[0]
}

func singleNumberOld(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return nums[0]
	}
	sort.Ints(nums)

	for i := 0; i < size-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[size-1]
}

func main() {
	nums := []int{2, 2, 1}

	fmt.Println(singleNumber(nums))
}
