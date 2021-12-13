package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func intersect(arr1 []int, arr2 []int) []int {
	hMap1 := make(map[int]int, len(arr1))
	res := make([]int, 0)

	for _, v := range arr1 {
		hMap1[v]++
	}

	for _, v := range arr2 {
		if data, _ := hMap1[v]; data > 0 {
			hMap1[v]--
			res = append(res, v)
		}
	}
	return res
}

func intersectOld(arr1 []int, arr2 []int) []int {
	hMap1 := make(map[int]int, len(arr1))
	hMap2 := make(map[int]int, len(arr2))
	for _, v := range arr1 {
		hMap1[v]++
	}
	for _, v := range arr2 {
		hMap2[v]++
	}

	var res = make([]int, 0)
	if len(hMap1) > len(hMap2) {
		hMap1, hMap2 = hMap2, hMap1
	}
	for v1, c1 := range hMap1 {
		if c2, ok := hMap2[v1]; ok {
			for i := 0; i < min(c1, c2); i++ {
				res = append(res, v1)
			}
		}
	}

	return res
}

func main() {
	arr1 := []int{1, 2, 2, 1}
	arr2 := []int{2, 2}

	fmt.Println(intersect(arr1, arr2))
}
