package main

import "fmt"

// func permute(n []int) [][]int {
// 	if len(n) == 1 {
// 		return [][]int{n}
// 	}

// 	res := [][]int{}

// 	for j := 0; j < len(n); j++ {
// 		// fmt.Println(res)
// 		first := n[0]
// 		n = n[1:]

// 		perms := permute(n)
// 		// fmt.Println(perms)
// 		// fmt.Println()

// 		for i := 0; i < len(perms); i++ {
// 			perms[i] = append(perms[i], first)
// 			// fmt.Println(perms[i])
// 		}
// 		fmt.Println(perms)

// 		res = append(res, perms...)
// 		n = append(n, first)
// 	}

// 	return res
// }

func permute(n []int) [][]int {
	var res = [][]int{}
	helper(&res, n, 0)
	return res
}

func helper(res *[][]int, n []int, currId int) {
	// fmt.Println("currArr:", currArr)

	if currId >= len(n) {
		t := make([]int, len(n))
		copy(t, n)
		*res = append(*res, t)
		return
	}

	for i := currId; i < len(n); i++ {
		n[currId], n[i] = n[i], n[currId]
		helper(res, n, currId+1)
		n[currId], n[i] = n[i], n[currId]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(len(permute(nums)))
}
