package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	res := ""

	hMap := make(map[int]byte)
	for i := 0; i < len(indices); i++ {
		hMap[indices[i]] = s[i]
	}

	for i := 0; i < len(indices); i++ {
		res += string(hMap[i])
		fmt.Println(res, i)
	}

	return res
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	fmt.Println(restoreString(s, indices))
}
