package main

import "fmt"

func rotateOld(nums []int, k int) {
	size := len(nums)
	for i := 0; i < k; i++ {
		temp := nums[size-1]
		copy(nums[1:], nums[:size-1])
		nums[0] = temp
	}
}

func rotateOld2(nums []int, k int) {
	size := len(nums)
	k %= size
	temp := make([]int, k)
	copy(temp[:], nums[size-k:])
	copy(nums[k:], nums[:size-k])
	copy(nums[:k], temp[:])
}

func rotate(nums []int, k int) {
	k = len(nums) - (k % len(nums))
	copy(nums[:], append(nums[k:], nums[:k]...))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(nums)
	rotate(nums, k)
	fmt.Println(nums)
}
