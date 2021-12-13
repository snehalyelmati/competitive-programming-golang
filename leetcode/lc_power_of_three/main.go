package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	num := math.Log(float64(n)) / math.Log(3)
	// isWhole := int(num*math.MaxInt32) == int(num)*math.MaxInt32
	isWhole := float64(int64(num)) == num
	fmt.Println(num, isWhole)
	return n > 0 && isWhole
}

func main() {
	n := 19684
	fmt.Println(isPowerOfThree(n))
}
