package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 || nums[nums[i]-1] == nums[i] {
			i++
		} else {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	res := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if i != nums[i-1] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	n := []int{1, 2}
	// n := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(n))
}

// 3, 2, 3, 4, 8, 2, 7, 1
// 1, 2, 3, 4, 5, 6, 7, 8
