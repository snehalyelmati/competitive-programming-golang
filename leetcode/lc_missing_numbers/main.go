package main

import (
	"fmt"
	"sort"
)

func missingNumberOld(n []int) int {
	sort.Ints(n)
	fmt.Println(n)
	for i, v := range n {
		if i != v {
			return i
		}
	}
	return len(n)
}

func missingNumberOld2(n []int) int {
	hMap := make(map[int]int)
	for _, v := range n {
		hMap[v] = v
	}
	for i := 0; i < len(n); i++ {
		if _, ok := hMap[i]; !ok {
			return i
		}
	}
	return len(n)
}

func missingNumberOld3(n []int) int {
	miss := len(n)
	for i, v := range n {
		miss ^= i ^ v
		fmt.Println(miss)
	}
	return miss
}

func missingNumberOld4(n []int) int {
	size := len(n)
	exp := size * (size + 1) / 2
	sum := 0
	for _, v := range n {
		sum += v
	}
	return exp - sum
}

func missingNumber(nums []int) int {
	mark := -1
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			mark = i
			i++
		} else if nums[i] == i+1 {
			i++
		} else {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	return mark + 1
}

func main() {
	n := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(n))
}
