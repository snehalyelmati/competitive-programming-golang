package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	fmt.Println(nums)
	fmt.Println("Size: ", len(nums))
	pivot := findPivot(nums)
	fmt.Println("Pivot: ", pivot)
    if pivot == -1 {
        return binarySearch(nums, target)
    }

	idx := binarySearch(nums[:pivot], target)
	if idx == -1 && len(nums[pivot:]) != 0 {
		idx = binarySearch(nums[pivot:], target)
		if idx != -1 {
			return pivot + idx
		}
	}
	return idx
}

// returns starting index of the second array
func findPivot(nums []int) int {
	l, h := 0, len(nums)-1

	for i := l; l < h; i++ {
		mid := (l + h) / 2
		fmt.Println(l, mid, h)
		fmt.Println("Nums: ", nums[l], nums[mid], nums[h])
		if mid > l && nums[mid] < nums[mid-1] {
			return mid
		}
		if mid < h && nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[l] > nums[mid] {
			// check left
			h = mid - 1
		} else {
			// check right
			l = mid + 1
		}
		fmt.Println(l, mid, h)
		fmt.Println("Nums: ", nums[l], nums[mid], nums[h])
	}

	return -1
}

func binarySearch(nums []int, target int) int {
	fmt.Println(nums)
	l, h := 0, len(nums)-1
	for i := l; l < h; i++ {
		mid := (l + h) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Println(nums[l], target)
	if nums[l] != target {
		return -1
	}
	return l
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(arr, target))
}
