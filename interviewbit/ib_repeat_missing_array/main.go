package main

import (
	"fmt"
	"sort"
)

/**
* @input A : Integer array
*
* @Output Integer array.
 */
func repeatedNumber(nums []int) []int {
	sort.Ints(nums)
	res := []int{0, 0} // repeated, missing
	for i := range nums {
		if nums[i] == nums[i+1] {
			res[0] = nums[i]
			nums[i+1] = 0
			break
		}
	}

	sum := 0
	for _, d := range nums {
		sum += d
	}
	n := len(nums)
	expSum := n * (n + 1) / 2
	res[1] = expSum - sum

	return res
}

func main() {
	nums := []int{3, 1, 2, 5, 3}
	fmt.Println(repeatedNumber(nums))
}
