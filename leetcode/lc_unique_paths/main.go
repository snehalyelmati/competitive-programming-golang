package main

import "fmt"

func uniquePaths(m int, n int) int {
	memory := make([][]int, m)
	for i := range memory {
		memory[i] = make([]int, n)
		for j := range memory[i] {
			memory[i][j] = -1
		}
	}
	return helper(m, n, 0, 0, memory)
}

func helper(m, n, i, j int, memory [][]int) int {
	if i >= m || j >= n {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	if memory[i][j] != -1 {
		return memory[i][j]
	} else {
		memory[i][j] = helper(m, n, i+1, j, memory) + helper(m, n, i, j+1, memory)
		return memory[i][j]
	}
}

func main() {
	m, n := 3, 7
	fmt.Println(uniquePaths(m, n))
}
