package main

import (
	"fmt"
	"sort"
)

func quickSort(nums []int) []int {
	helper(&nums, 0, len(nums)-1)
	return nums
}

func helper(nums *[]int, l, h int) {
	if h-l <= 0 {
		return
	}

	p := (l + h) / 2
	s, e := l, h
	// fmt.Println(s, e)
	// fmt.Println("Pivot: ", (*nums)[p])

	// prepPivot
	for s < e {
		// fmt.Println((*nums)[s], (*nums)[e])
		if (*nums)[s] >= (*nums)[p] && (*nums)[e] <= (*nums)[p] {
			if p == s {
				p = e
			} else if p == e {
				p = s
			}

			(*nums)[s], (*nums)[e] = (*nums)[e], (*nums)[s]
			s++
			e--
		}
		if (*nums)[s] < (*nums)[p] {
			s++
		}
		if (*nums)[e] > (*nums)[p] {
			e--
		}
	}
	// fmt.Println(nums, p, (*nums)[p])

	// repeat this for both halves of the arrays
	helper(nums, l, p-1)
	helper(nums, p+1, h)
}

func main() {
	nums := []int{5, 3, 6, 4, 1, 2}
	fmt.Println(quickSort(nums))
	sort.Ints(nums)
	fmt.Println(nums)
}

// 5, 3, 6, 4, 1, 2
// 1, 2, 3, 4, 6, 5
