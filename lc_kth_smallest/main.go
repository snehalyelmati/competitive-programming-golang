package main

import "fmt"

func kthSmallest(m [][]int, k int) int {
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
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8

	fmt.Println(kthSmallest(matrix, k))
}
