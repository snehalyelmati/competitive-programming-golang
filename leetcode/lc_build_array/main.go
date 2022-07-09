package main

import "fmt"

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums))
}
