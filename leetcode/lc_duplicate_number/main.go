package main

import (
	"fmt"
	"sort"
)

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// fmt.Printf("slow: %v %v\n", slow, nums[slow])
	// fmt.Printf("fast: %v %v\n", fast, nums[fast])

	walker := 0
	for walker != slow {
		slow = nums[slow]
		walker = nums[walker]
	}
	// fmt.Printf("slow: %v %v\n", slow, nums[slow])
	// fmt.Printf("walker: %v %v\n", walker, nums[walker])

	return walker
}

func findDuplicateOld4(nums []int) int {
	sumNums := 0
	sumId := 0
	maxNums := -1

	for i, v := range nums {
		if maxNums < v {
			maxNums = v
		}
		sumNums += v
		sumId += i
	}

	rem := sumNums - sumId
	lenDiff := len(nums) - maxNums

	return rem/lenDiff
}

func findDuplicateOld3(n []int) int {
	sort.Ints(n)
	for i := 1; i < len(n); i++ {
		if n[i-1] == n[i] {
			return n[i]
		}
	}
	return -1
}

func findDuplicateOld2(nums []int) int {
	hMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			return v
		}
		hMap[v] = true
	}
	return -1
}

func findDuplicateOld(nums []int) int {
	sumNums := 0
	sumId := 0

	for i, v := range nums {
		sumNums += v
		sumId += i
	}

	return sumNums - sumId
}

func main() {
	nums := []int {1, 3, 4, 2, 2}
	fmt.Printf("nums: %v\n", nums)
	fmt.Println(findDuplicate(nums))
}