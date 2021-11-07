package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	hMap := make(map[int]int)
	for _, v := range nums2 {
		hMap[v]++
	}

	res := []int{}
	for _, v := range nums1 {
		if data, ok := hMap[v]; ok && data > 0 {
			res = append(res, v)
			hMap[v]--
		}
	}
	return res
}

func main() {
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2}
	fmt.Println(intersect(nums1, nums2))
}
