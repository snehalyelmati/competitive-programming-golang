package main

import "fmt"

func rotateOld(nums []int, k int) {
	size := len(nums)
	for i := 0; i < k; i++ {
		temp := nums[size-1]
		copy(nums[1:], nums[:size-1])
		// fmt.Println(nums)
		nums[0] = temp
	}
}

func rotate(nums []int, k int) {
	k = (2*len(nums) - k) % len(nums)
	fmt.Println(k, nums[k:], nums[:k])
	copy(nums[:], append(nums[k:], nums[:k]...))
	fmt.Println(nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(nums)
	rotate(nums, k)
	fmt.Println(nums)
}
