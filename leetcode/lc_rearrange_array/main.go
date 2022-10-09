package main

import "fmt"

// find the anomaly
// find the wanted element - either positive or negative
// either swap it back or build a mod array with the wanted element in the right place
func rearrangeArray(nums []int) []int {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 && nums[i] < 0 {
			temp := i
			for nums[temp] < 0 {
				fmt.Println(temp)
				temp++
			}
			fmt.Println(nums[i], nums[temp])
			nums = moveEle(nums, temp, i)
		} else if i%2 != 0 && nums[i] > 0 {
			temp := i
			for nums[temp] > 0 {
				fmt.Println(temp)
				temp++
			}
			fmt.Println(nums[i], nums[temp])
			nums = moveEle(nums, temp, i)
		}
		fmt.Println(nums)
	}

	return nums
}

func moveEle(nums []int, temp, pos int) []int {
	fmt.Printf("Moving element from %d to %d, elements are %d, %d\n", temp, pos, nums[temp], nums[pos])
	for ; temp > pos; temp-- {
		nums[temp], nums[temp-1] = nums[temp-1], nums[temp]
		fmt.Println("-", nums, temp, temp-1)
	}
	return nums
}

// finds the first index of positive and negative numbers where the condition breaks
func set(nums []int, p, n int) (int, int) {
	for p < len(nums) && nums[p] > 0 {
		p += 2
	}
	for n < len(nums) && nums[n] < 0 {
		n += 2
	}
	return p, n
}

func main() {
	// nums := []int{3, 1, -2, -5, 2, -4}
	nums := []int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}
	fmt.Println(rearrangeArray(nums))
}
