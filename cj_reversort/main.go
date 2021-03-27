package main

import (
	"fmt"
	"strconv"
)

func reverseSort(arr []int) int {
	// find the index of the minimum element --> j
	// reverse the list from i to j, where i is the current index

	count := 0
	for j := 0; j < len(arr)-1; j++ {
		min := j
		// for i, v := range arr[j:] {
		// 	if v < arr[min] {
		// 		min = i
		// 	}
		// }

		for i := j; i < len(arr); i++ {
			if arr[min] > arr[i] {
				min = i
			}
		}

		count += min - j + 1

		// fmt.Println(min - j + 1)
		// fmt.Println(arr[j:])
		// fmt.Println("Min index:", min, arr[min], j)
		// fmt.Println(arr)
		for i, k := j, min; i < k; i, k = i+1, k-1 {
			arr[i], arr[k] = arr[k], arr[i]
		}
		// fmt.Println(arr)
	}

	return count
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 1; t <= T; t++ {
		var N int
		fmt.Scan(&N)

		arr := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&arr[i])
		}

		// fmt.Println(N, arr)
		count := reverseSort(arr)
		fmt.Println("Case #"+strconv.Itoa(t)+":", count)
	}
}
