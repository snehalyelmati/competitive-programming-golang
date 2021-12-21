package main

import "fmt"

func solve(A []int, B int) int {
	hMap := make(map[int]int)

	xor := 0
	count := 0
	for _, d := range A {
		xor ^= d
		fmt.Println(xor)
		if xor == B {
			count++
		}
		if freq, ok := hMap[xor^B]; ok {
			count += freq
		}
		hMap[xor]++
	}
	fmt.Println(hMap)

	return count
}

func main() {
	nums := []int{5, 6, 7, 8, 9}
	target := 5

	fmt.Println(solve(nums, target))
}
