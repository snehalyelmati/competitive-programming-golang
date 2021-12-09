package main

import "fmt"

func sortColors(nums []int) {
	i := 0
	j := len(nums) - 1

	for i <= j {
		if nums[i] == 0 {
			// put it in front of the array
			copy(nums[1:], append(nums[:i], nums[i+1:]...))
			nums[0] = 0
			i++
		} else if nums[i] == 1 {
			// do nothing - increment
			i++
		} else if nums[i] == 2 {
			// put it at the end of the array
			copy(nums[:len(nums)-1], append(nums[:i], nums[i+1:]...))
			nums[len(nums)-1] = 2
			j--
		}
	}

	fmt.Println("Sorted nums: ", nums)
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}
