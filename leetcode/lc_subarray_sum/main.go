package main

import "fmt"

// this works for non-negative integers
func subarraySumNonNegative(nums []int, k int) int {
	count, sum := 0, 0
	l, h := -1, -1

	for l < len(nums) && h < len(nums) {
		fmt.Println(h, l, sum)
		if sum <= k {
			if sum == k && h != l {
				count++
				fmt.Println(count)
			}

			h++
			if h >= len(nums) {
				break
			}
			sum += nums[h]
		} else if sum > k {
			l++
			sum -= nums[l]
		}
	}

	return count
}

func subarraySum(nums []int, k int) int {
	hMap := make(map[int]int)
	hMap[0] = 1
	currSum, count := 0, 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		fmt.Printf("Current sum: %d, diff: %d \n", currSum, currSum-k)
		if d, ok := hMap[currSum-k]; ok {
			count += d
			fmt.Println("Increasing count by", d)
		}
		hMap[currSum]++
	}
	fmt.Println(hMap)

	return count
}

func main() {
	// nums := []int{1, 10, 1, 1}
	// nums := []int{1, 2, 3}
	nums := []int{1, -1, 0}
	// currSum - k = x
	k := 0
	fmt.Printf("Array: %d, target: %d\n", nums, k)
	fmt.Println("Subarray sum:", subarraySum(nums, k))
}
