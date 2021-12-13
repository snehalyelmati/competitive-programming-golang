package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0

	for i := 0; i < len(points)-1; i++ {
		x := math.Abs(float64(points[i][0] - points[i+1][0]))
		y := math.Abs(float64(points[i][1] - points[i+1][1]))
		sum += max(x, y)
		fmt.Println(sum, x, y)
	}

	return sum
}

func max(x, y float64) int {
	if x > y {
		return int(x)
	}
	return int(y)
}

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println(minTimeToVisitAllPoints(points))
}
