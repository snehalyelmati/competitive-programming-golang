package main

import "fmt"

func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			// nums = append(append(nums[:i], nums[i+1:]...), 0)
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
		}
		fmt.Println(nums)
	}
}

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 1, 0, 3, 12}

	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}

