package main

import (
	"fmt"
	"math"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	// find max height and max width - max consequtive difference
	// multiply it

	// might not be needed
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	// basic refractoring
	horizontalCuts = append([]int{0}, horizontalCuts...)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append([]int{0}, verticalCuts...)
	verticalCuts = append(verticalCuts, w)

	maxH := 0
	for i:=1; i<len(horizontalCuts); i++ {
		diff := horizontalCuts[i]-horizontalCuts[i-1]
		if diff > maxH {
			maxH = diff
		}
	}

	maxW := 0
	for i:=1; i<len(verticalCuts); i++ {
		diff := verticalCuts[i]-verticalCuts[i-1]
		if diff > maxW {
			maxW = diff
		}
	}
	fmt.Println("Max height and width:", maxH, maxW)

	return maxH*maxW%(int(math.Pow10(9)+7))
}

func main() {
	h, w := 5, 4
	// horizontalCuts := []int{1,2,4}
	// verticalCuts := []int{1,3}

	horizontalCuts := []int{3,1}
	verticalCuts := []int{1}

	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts))
}
