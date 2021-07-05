package main

import (
	"fmt"
)

// basically implement binary search for array:
// - find the mid element
// - calculate the elements less than that number in the array
// - based on that select the part of the flow based on the condition > or < mid
func kthSmallest(m [][]int, k int) int {
	n := len(m)
	l := m[0][0]
	h := m[n-1][n-1]
	mid := (l + h) / 2

	for l < h {
		fmt.Printf("l: %v\n", l)
		fmt.Printf("mid: %v\n", mid)
		fmt.Printf("h: %v\n", h)

		numbersLesser := calculateNumbersLesser(m, mid, n)
		fmt.Printf("numbersLesser: %v\n", numbersLesser)

		if numbersLesser < k {
			l = mid + 1
		} else {
			// if k == numbersLesser
			// right endpoint can be revised
			h = mid
		}
		mid = (l + h) / 2
		fmt.Println()
	}

	return l
}

// calculate the number of elements less than x
func calculateNumbersLesser(m [][]int, x, n int) int {
	count := 0
	i := 0
	j := n - 1

	fmt.Printf("x: %v\n", x)

	for i >= 0 && j >= 0 && i < n && j < n {
		fmt.Printf("m[i][j]: %v\n", m[i][j])
		if x >= m[i][j] {
			count += j + 1
			i++
		} else if j > 0 && x > m[i][j-1] {
			count += j
			i++
		} else {
			j--
		}
		fmt.Printf("count: %v\n", count)
	}

	return count
}

func kthSmallestOld(m [][]int, k int) int {
	res := make([]int, 0)
	fmt.Println("Size of the array is:", len(res))

	// TODO:
	// merge(arr, part_of_the_matrix)
	// return that part of the array

	for _, v := range m {
		fmt.Printf("res: %v\n", res)
		res = merge(res, v)
	}

	fmt.Println("Sorted array: ", res)

	return res[k-1]
}

func merge(a []int, b []int) []int {
	fmt.Println()

	aLen := 0
	bLen := 0

	res := make([]int, len(a)+len(b))
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	i := 0
	for i = 0; aLen < len(a) && bLen < len(b); i++ {
		if a[aLen] < b[bLen] {
			res[i] = a[aLen]
			// res = append(res, a[aLen])
			aLen++
		} else {
			res[i] = b[bLen]
			// res = append(res, b[bLen])
			bLen++
		}
	}

	fmt.Printf("res[i:]: %v\n", res[i:])
	if aLen != len(a) {
		fmt.Printf("a[aLen:]: %v\n", a[aLen:])
		copy(res[i:], a[aLen:])
		// res = append(res, a[aLen:]...)
	} else if bLen != len(b) {
		fmt.Printf("b[bLen:]: %v\n", b[bLen:])
		copy(res[i:], b[bLen:])
		// res = append(res, b[bLen:]...)
	}

	return res
}

func main() {
	// matrix := [][]int{
	// 	{1, 5, 9},
	// 	{10, 11, 13},
	// 	{12, 13, 15},
	// }
	matrix := [][]int{
		{1, 2},
		{1, 3},
	}

	k := 1

	fmt.Println(kthSmallest(matrix, k))
	// fmt.Println(calculateNumbersLesser(matrix, (1+15)/2, 3))
}
