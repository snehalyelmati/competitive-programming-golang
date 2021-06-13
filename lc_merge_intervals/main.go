package main

import (
	"fmt"
	"math"
	"sort"
)

func merge(n [][]int) [][]int {
	sort.Slice(n, func(i, j int) bool {
		return n[i][0] < n[j][0]
	})
	
	res := make([][]int, 0)
	res = append(res, n[0])
	r := 0

	for i := 1; i < len(n); i++ {
		if res[r][1] >= n[i][0] {
			// merge
			res[r][1] = int(math.Max(float64(res[r][1]), float64(n[i][1])))
		} else {
			res = append(res, n[i])
			r++
		}
	}
	fmt.Println(res)

	return res
}

func mergeOld2(n [][]int) [][]int {
	sort.Slice(n, func(i, j int) bool {
		return n[i][0] < n[j][0]
	})
	
	for i:=0; i < len(n)-1; i++ {
		if n[i][1] >= n[i+1][0] {
			n[i][1] = int(math.Max(float64(n[i][1]), float64(n[i+1][1])))
			n = append(n[:i+1], n[i+2:]...)
			i -= 1
		}
	}

	return n
}

func mergeOld(n [][]int) [][]int {
	sort.Slice(n, func(i, j int) bool {
		return n[i][0] < n[j][0]
	})
	
	for i:=0; i < len(n)-1; i++ {
		if n[i][1] >= n[i+1][0] {
			if n[i][1] < n[i+1][1] {
				n[i][1] = n[i+1][1]
			} else {
				n[i+1][1] = n[i][1]
			}
			n = append(n[:i+1], n[i+2:]...)
			i -= 1
		}
	}

	return n
}

func main() {
	nums := [][]int {
		// {1, 4},
		{5, 6},
		{1, 4},
		{2, 3},
		// {8, 10},
		// {15, 18},
	}

	fmt.Println(merge(nums))
}