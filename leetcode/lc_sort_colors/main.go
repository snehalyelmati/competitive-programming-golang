package main

import "fmt"

func sortColors(nums []int) {
	zero, two := 0, len(nums)-1
	for i := 0; i <= two && two >= 0 && zero < len(nums); {
		fmt.Println(nums, i, zero, two)
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			for zero < len(nums) && nums[zero] == 0 {
				zero++
				i++
			}
		} else if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			for two >= 0 && nums[two] == 2 {
				two--
			}
		} else if nums[i] == 1 {
			i++
		}

	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println()
	fmt.Println(nums)
}
