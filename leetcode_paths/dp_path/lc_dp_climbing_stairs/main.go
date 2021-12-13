package main

import "fmt"

func climbStairs(n int) int {
	memory := make([]int, n)
	return helper(n, memory)
}

func helper(n int, memory []int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if memory[n-1] != 0 {
		return memory[n-1]
	}

	memory[n-1] = helper(n-1, memory) + helper(n-2, memory)
	return memory[n-1]
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
