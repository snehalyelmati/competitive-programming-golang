package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	for i, j, k := 0, len(nums)-1, len(nums)-1; i <= j; {
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return res
}

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	nums := []int{-7,-3,2,3,11}
	fmt.Println(sortedSquares(nums))
}
