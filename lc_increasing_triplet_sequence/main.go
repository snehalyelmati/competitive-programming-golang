package main

import (
	"fmt"
	"math"
)

func increasingTriplet(arr []int) bool {
	i := math.MaxInt32
	j := math.MaxInt32

	for _, v := range arr {
		if v <= i {
			i = v
		} else if v <= j {
			j = v
		} else {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(increasingTriplet(arr))
}
