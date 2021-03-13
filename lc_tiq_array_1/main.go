package main

import "fmt"

func shiftArr(arr []int) []int {
	for len(arr) > 1 && arr[0] == arr[1] {
		arr = shiftArr(arr[1:])
	}
	return arr
}

func removeDuplicatesOld(nums []int) int {
	var size = len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return 1
	}
	for i := 0; i < size; i++ {
		temp := shiftArr(nums[i:])
		copy(nums[i:], temp)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return i
		}
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	u := 0
	for i := 1; i < size; i++ {
		if nums[i] != nums[u] {
			u++
			nums[u] = nums[i]
		}
	}
	return u + 1
}

func main() {
	var nums = []int{0, 1, 1, 1, 2, 2, 3, 3, 4}

	res := removeDuplicates(nums[:])
	fmt.Println(nums, res)
}
