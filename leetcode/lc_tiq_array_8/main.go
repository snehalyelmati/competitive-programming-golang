package main

import "fmt"

func moveZeroes(nums []int) {
	size := len(nums)
	for i := size - 1; i >= 0; i-- {
		if nums[i] == 0 {
			copy(nums[i:], nums[i+1:])
			nums[size-1] = 0
		}
	}
}

func moveZeroesOld(nums []int) {
	res := make([]int, 0, len(nums))
	zeroes := make([]int, 0, len(nums))

	for _, v := range nums {
		if v == 0 {
			zeroes = append(zeroes, 0)
		} else {
			res = append(res, v)
		}
	}
	copy(nums[:len(res)], res[:])
	copy(nums[len(res):], zeroes[:])
	// res = append(res, zeroes...)
	// fmt.Println(res)
	// nums = res
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
