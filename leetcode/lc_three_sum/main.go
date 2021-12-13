package main

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	ans := make(map[string]bool)
	res := make([][]int, 0)
	hMap := make(map[int]int)
	for j, v1 := range nums {
		hMap[v1] = j
	}
	for v, i := range hMap {
		for j := i + 1; j < len(nums); j++ {
			// fmt.Println(v, nums[j], hMap[-v-nums[j]])
			// fmt.Println(nums[j:])
			if _, ok := hMap[-v-nums[j]]; ok && i != j && hMap[-v-nums[j]] != i && j != hMap[-v-nums[j]] {
				temp := []int{v, nums[j], -v - nums[j]}
				// fmt.Println(temp, i, j, hMap[-v-nums[j]])
				sort.Ints(temp)
				str := strconv.Itoa(temp[0]) + ", " + strconv.Itoa(temp[1]) + ", " + strconv.Itoa(temp[2])
				// fmt.Println(str)
				if _, present := ans[str]; !present {
					res = append(res, []int{temp[0], temp[1], temp[2]})
					ans[str] = true
				}
			}
			// fmt.Println(ans)
		}
	}

	return res
}

func main() {
	arr := []int{3, -2, 1, 0}
	fmt.Println(threeSum(arr))
}
