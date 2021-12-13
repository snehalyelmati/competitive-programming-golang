package main

import (
	"fmt"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func rob(n []int) int {
	if len(n) == 1 {
		return n[0]
	} else if len(n) == 2 {
		return max(n[0], n[1])
	}

	dp := make([]int, len(n))

	dp[0] = n[0]
	dp[1] = max(n[0], n[1])

	for i := 2; i < len(n); i++ {
		dp[i] = max(dp[i-2]+n[i], dp[i-1])
	}
	return dp[len(n)-1]
}

func main() {
	arr := []int{1}
	fmt.Println(rob(arr))
}
