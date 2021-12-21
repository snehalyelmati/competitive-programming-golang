package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); {
		temp := 0
		for j := i + 1; j < len(nums); {
			left := j + 1
			right := len(nums) - 1
			subTarget := target - nums[i] - nums[j]

			for left < right {
				if nums[left]+nums[right] == subTarget {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					// skip left
					temp = left
					for left < len(nums) && nums[temp] == nums[left] {
						left++
					}

					// skip right
					temp = right
					for right > 0 && nums[temp] == nums[right] {
						right--
					}
				} else if nums[left]+nums[right] >= subTarget {
					// skip right
					temp = right
					for right > 0 && nums[temp] == nums[right] {
						right--
					}
				} else {
					// skip left
					temp = left
					for left < len(nums) && nums[temp] == nums[left] {
						left++
					}
				}
			}

			// skip j duplicates
			temp = j
			for j < len(nums) && nums[temp] == nums[j] {
				j++
			}
		}

		// skip i duplicates
		temp = i
		for i < len(nums) && nums[temp] == nums[i] {
			i++
		}
	}

	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(target)
	fmt.Println(fourSum(nums, target))
}
