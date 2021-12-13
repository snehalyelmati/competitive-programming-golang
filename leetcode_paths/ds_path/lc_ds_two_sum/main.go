package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, ele := range nums {
		if data, ok := hashMap[target-ele]; ok {
			return []int{i, data}
		}
		hashMap[ele] = i
	}
	return []int{0, 0}
}

func main() {
	nums := []int{2,7,11,15}
	target := 18
	fmt.Println("Indices: ", twoSum(nums, target))
}
