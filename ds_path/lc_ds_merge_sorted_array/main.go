package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	i := 0
	j := 0
	k := 0
	for i = 0; i < m+n && j < m && k < n; i++ {
		if nums1[j] >= nums2[k] {
			res[i] = nums2[k]
			k++
		} else {
			res[i] = nums1[j]
			j++
		}
	}

	if k == n && i != m+n {
		copy(res[i:], nums1[j:m])
	}
	if j == m && i != m+n {
		copy(res[i:], nums2[k:n])
	}
	copy(nums1, res)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println(nums1)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
