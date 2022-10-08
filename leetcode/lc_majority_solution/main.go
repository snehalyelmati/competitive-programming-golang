package main

import (
	"fmt"
	"sort"
)

func majorityElementOld(n []int) int {
	hMap := make(map[int]int)
	criteria := len(n) / 2

	for _, v := range n {
		if _, ok := hMap[v]; !ok {
			hMap[v] = 0
		} else {
			hMap[v]++
			if hMap[v] >= criteria {
				return v
			}
		}
	}
	// for k, v := range hMap {
	// 	if v >= criteria {
	// 		return k
	// 	}
	// }
	return -1
}

func majorityElement(n []int) int {
	sort.Ints(n)
	count := 1
	criteria := len(n) / 2

	for i := 0; i < len(n)-1; i++ {
		if n[i] != n[i+1] && count > criteria {
			return n[i]
		} else if n[i] != n[i+1] {
			count = 1
		} else {
			fmt.Println(n[i], count, criteria)
			count++
			if count > criteria {
				return n[i]
			}
		}
	}

	return n[0]
}

func majorityElementOptimized(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    candidate, count := nums[0], 1
    for i:=1; i<len(nums); i++ {
        if nums[i] == candidate {
            count++
        } else {
            count--
            if count == 0 {
                candidate = nums[i]
                count = 1
            }
        }
    }
    return candidate
}

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(arr))
}
