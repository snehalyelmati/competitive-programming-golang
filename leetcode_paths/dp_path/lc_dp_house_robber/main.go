package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	mem := make([]int, len(nums))

	mem[0] = nums[0]
	mem[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < len(nums); i++ {
		mem[i] = int(math.Max(float64(mem[i-2])+float64(nums[i]), float64(mem[i-1])))
	}

	return mem[len(nums)-1]
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	// nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}


// 0 1 2 3 4
//       *
