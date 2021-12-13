package main

import "fmt"

func twoSumOld(nums []int, target int) []int {
	var res = make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0], res[1] = i, j
				// fmt.Printf("[%d,%d]\n", nums[i], nums[j])
				break
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	// var res = make([]int, 2)
	siz := len(nums)
	var lookupMap = make(map[int]int, siz)
	for i, v := range nums {
		lookupMap[v] = i
	}
	// fmt.Println(lookupMap)

	for i, v := range nums {
		// fmt.Println(i, v, target-v, lookupMap[target-v])
		if data, ok := lookupMap[target-v]; ok {
			fmt.Println(i, data)
			if i==data {
				continue
			}
			return []int{i, data}
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(nums)

	target := 7 + 15
	// fmt.Println(target)

	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumOld(nums, target))
}
