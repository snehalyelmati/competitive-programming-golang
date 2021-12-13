package main

import (
	"fmt"
)

func mergeOld(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	j := 0
	i := 0
	// sort.Ints(nums1[:m])
	// sort.Ints(nums2[:])
	// fmt.Println(nums1, nums2)

	for i < len(nums2) && j < len(nums1) {
		// fmt.Println(nums1[j], nums2[i])
		if nums1[j] < nums2[i] && j <= m {
			fmt.Println(nums1[j+1:], nums1[j:])
			j++
		} else if nums1[j] < nums2[i] {
			// fmt.Println(nums1[j+1:], nums1[j:])
			j++
			nums1[j] = nums2[i]
			i++
		} else {
			copy(nums1[j+1:], nums1[j:])
			nums1[j] = nums2[i]
			fmt.Println(nums1, j, nums1[j+1:], nums1[j:])
			j += 1
			// fmt.Println(nums1, nums1[j], j, nums2[i], i)
			// if nums1[j] != nums2[i] {
			// 	j += 1
			// }
			i++
		}
		fmt.Println(nums1)
	}
	// copy(nums1[j+1:], nums2[i:])
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, 0)
	i := 0
	j := 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}

	if i == m {
		arr = append(arr, nums2[j:]...)
	}
	if j == n {
		arr = append(arr, nums1[i:]...)
	}
	copy(nums1, arr)
}

func main() {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1

	// fmt.Println(nums1)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
