package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
