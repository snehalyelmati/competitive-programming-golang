package main

import "fmt"

func firstMissingPositive(nums []int) int {
	// use cyclic sort
	for i := 0; i < len(nums); {
		fmt.Println(nums[i])
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] && nums[i] != i+1 {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	fmt.Println(nums)

	// identify the first missing number
	i := 1
	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	fmt.Println("hello")
	return len(nums) + 1
}

func main() {
	// nums := []int{1, 2, 0}
	// nums := []int{2, 1}
	nums := []int{3, 4, -1, 1}
	// nums := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums))
}
