package main

import (
	"fmt"
	"math"
)

func twoSumClosest(nums []int, k int) []int {
	// num - closest to it
	// hMap := make(map[int]int)
	// for _, d := range nums {
	// 	hMap[d] = -1
	// }

	min := math.MaxInt
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if min > k-nums[i]-nums[j] {
				min = k - nums[i] - nums[j]
				res = []int{nums[i], nums[j]}
			}
		}
	}

	return res
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// target := 10

	nums := []int{-1, 2, 1, -4}
	target := 4

	fmt.Println(twoSumClosest(nums, target))
}
