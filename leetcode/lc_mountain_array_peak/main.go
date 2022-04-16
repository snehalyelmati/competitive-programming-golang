package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	return binarySearch(arr)
}

func binarySearch(arr []int) int {
	l, h := 0, len(arr)
	mid := (l + h) / 2

	for i := 0; l != h; i++ {
		if arr[mid] > arr[mid+1] {
			h = mid
		} else {
			l = mid + 1
		}
		mid = (l + h) / 2
	}

	return mid
}

func main() {
	arr := []int{1, 2, 3, 4, 7, 6, 5, 4, 3}
	fmt.Println(peakIndexInMountainArray(arr))
}
