package main

import "fmt"

func splitArray(nums []int, m int) int {
	maxSum := 0
	minSum := 0
	for i := 0; i < len(nums); i++ {
		maxSum += nums[i]
		if nums[i] > minSum {
			minSum = nums[i]
		}
	}

	l, h := minSum, maxSum
	for i := l; l < h; i++ {
		mid := (l + h) / 2
		fmt.Println(mid)

		tempSum := 0
		count := 1
		for j := 0; j < len(nums); j++ {
			if tempSum+nums[j] > mid {
				count++
				tempSum = nums[j]
			} else {
				tempSum += nums[j]
			}
		}

		if count <= m {
			// if pieces are less than required => mid is more, decrease mid
			h = mid
		} else {
			// if pieces are more than required => mid is less, increase mid
			l = mid + 1
		}
	}

	return h
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2

	fmt.Println(splitArray(nums, m))
}
