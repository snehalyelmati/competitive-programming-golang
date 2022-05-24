package main

import "fmt"

func cyclicSort(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		}
		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
	}
	return nums
}

func main() {
	nums := []int{3, 5, 2, 1, 4}
	fmt.Println(cyclicSort(nums))
}
