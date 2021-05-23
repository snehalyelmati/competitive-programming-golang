package main

import "fmt"

var res = [][]int{}

func subsetsOld(nums []int) [][]int {
	helper(nums, []int{}, 0)
	return res
}

func helper(nums, curr []int, id int) {
	res = append(res, curr)

	for i := id; i < len(nums); i++ {
		helper(nums, append(curr, nums[i]), i+1)
	}
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return make([][]int, 1)
	}

	first := nums[len(nums)-1]
	restNums := nums[:len(nums)-1]

	combWithFirst := [][]int{}
	combWithoutFirst := subsets(restNums)
	for _, v := range combWithoutFirst {
		temp := append([]int{first}, v...)
		combWithFirst = append(combWithFirst, temp)
	}

	return append(combWithoutFirst, combWithFirst...)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
