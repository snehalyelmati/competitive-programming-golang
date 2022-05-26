package main

import "fmt"

func fibonacci(n int) int {
	arr := make([]int, n+1)
	arr[1] = 1
	return helper(&arr, n)
}

func helper(arr *[]int, n int) int {
	if n < 2 || (*arr)[n] != 0 {
		return (*arr)[n]
	}
	(*arr)[n] = helper(arr, n-1) + helper(arr, n-2)
	fmt.Println(arr)
	return (*arr)[n]
}

func binarySearch(arr []int, n int) int {
	return binarySearchHelper(arr, 0, len(arr), n)
}

func binarySearchHelper(arr []int, l, h, n int) int {
	mid := (l + h) / 2
	if arr[mid] == n {
		return mid
	} else if arr[mid] > n {
		return binarySearchHelper(arr, l, mid, n)
	} else if arr[mid] > n {
		return binarySearchHelper(arr, mid, h, n)
	} else {
		return -1
	}
}

func main() {
	// fmt.Println(fibonacci(50))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))
}
