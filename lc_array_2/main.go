package main

import "fmt"

func sum(nums []int) int {
	var res int = 0
	for _, v := range nums {
		// fmt.Println(v)
		res += v
	}
	// fmt.Println("Sum from func:", res)
	return res
}

func maxSubArrayOld(nums []int) int {
	var max int = nums[0]
	var sliSum int = 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sliSum = sum(nums[i : j+1])
			// fmt.Println(sliSum, max)
			if sliSum > max {
				max = sliSum
			}
			// fmt.Println("Sum: ", sliSum, ",", i, j)
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1]+nums[i] {
			nums[i] += nums[i-1]
		}
	}
	// to find max
	var max int = nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//-2, 1, -3, 4, 3,  5, 6, 1, 5
	fmt.Println(nums)

	fmt.Println(maxSubArrayOld(nums))
	fmt.Println(maxSubArray(nums))
}
