package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairsOld(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])))
	}
	return int(math.Min(float64(cost[len(cost)-1]), float64(cost[len(cost)-2])))
}

func minCostClimbingStairs(cost []int) int {
	mem := make([]int, len(cost)+1)
	return int(math.Min(float64(helper(cost, 0, mem)), float64(helper(cost, 1, mem))))
}

func helper(cost []int, idx int, mem []int) int {
	if idx >= len(cost) {
		return 0
	}
	if mem[idx] != 0 {
		return mem[idx]
	}
	step1 := helper(cost, idx+1, mem)
	step2 := helper(cost, idx+2, mem)
	mem[idx] = cost[idx] + int(math.Min(float64(step1), float64(step2)))
	return mem[idx]
}

func main() {
	// cost := []int{10, 15, 20}
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
