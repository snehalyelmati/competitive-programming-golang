package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println(temp)
		if temp+nums[i] < nums[i] {
			temp = nums[i]
		} else {
			temp += nums[i]
		}

		if max < temp {
			max = temp
		}
	}

	return max
}

func main() {
	num1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	num2 := []int{1, 2}
	fmt.Println(maxSubArray(num1))
	fmt.Println(maxSubArray(num2))
}
