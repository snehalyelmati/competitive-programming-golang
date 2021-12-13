package main

import "fmt"

func cs(n int, his map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if d, ok := his[n]; ok {
		return d
	}

	value := cs(n-1, his) + cs(n-2, his)
	his[n] = value

	return value
}

func climbStairs(n int) int {
	his := make(map[int]int)

	return cs(n, his)
}

func main() {
	n := 45

	fmt.Println(climbStairs(n))
}
