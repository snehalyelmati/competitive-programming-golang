package main

import "fmt"

func bubbleSort(nums []int) []int {
	helper(&nums, len(nums)-1, 0)
	return nums
}

func helper(nums *[]int, r, c int) {
	if r < 1 {
		return
	} else if c >= r {
		helper(nums, r-1, 0)
	} else {
		if (*nums)[c] > (*nums)[c+1] {
			(*nums)[c], (*nums)[c+1] = (*nums)[c+1], (*nums)[c]
		}
		helper(nums, r, c+1)
	}
}

func main() {
	nums := []int{}
	fmt.Println(bubbleSort(nums))
}

// 2, 3, 5, 1, 4
// 2, 3, 5, 1, 4
// 2, 3, 5, 1, 4
// 2, 3, 5, 1, 4
