package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min := math.MaxInt32
	maxProf := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > maxProf {
			maxProf = v - min
		}
	}

	return maxProf
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
