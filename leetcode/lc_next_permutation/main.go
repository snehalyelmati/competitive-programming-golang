package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	targetIdx := 0
	flag := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			targetIdx = i - 1
			flag = true
			break
		}
	}
	fmt.Println("Target idx: ", targetIdx)
	fmt.Println(nums)

	swapTarget := pickNextLargest(nums, targetIdx)
	fmt.Println("Swap idx: ", swapTarget)
	nums[swapTarget], nums[targetIdx] = nums[targetIdx], nums[swapTarget]
	fmt.Println(nums)

	if flag {
		sort.Ints(nums[targetIdx+1:])
	} else {
		sort.Ints(nums)
	}
}

func pickNextLargest(nums []int, targetIdx int) int {
	swapTarget := targetIdx + 1
	for i := targetIdx + 1; i < len(nums); i++ {
		fmt.Println(i, targetIdx, swapTarget)
		fmt.Println(nums[i], nums[targetIdx], nums[swapTarget])
		if nums[i] > nums[targetIdx] && nums[swapTarget] > nums[i] {
			swapTarget = i
		}
	}

	return swapTarget
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(nums)
	// fmt.Println(pickNextLargest(nums, 2))
	nextPermutation(nums)
	fmt.Println(nums)
}
