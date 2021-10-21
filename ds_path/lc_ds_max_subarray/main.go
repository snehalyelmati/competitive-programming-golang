package main

import (
	"fmt"
)

// return max subarray sum
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum += nums[i]
		if currSum < nums[i] {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{5, 4, -1, 7, 8}
	// nums := []int{-2, 1}
	fmt.Println("Max sum: ", maxSubArray(nums))
}
