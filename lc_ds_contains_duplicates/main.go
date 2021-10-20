package main

import (
	"fmt"
	"sort"
)

func containsDuplicateOld(nums []int) bool {
	sort.Ints(nums)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i]==nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	hMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			return true
		}
		hMap[v] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 1, 4}
	fmt.Printf("This array: %v, contains duplicates: %v.\n", nums, containsDuplicate(nums))
}
