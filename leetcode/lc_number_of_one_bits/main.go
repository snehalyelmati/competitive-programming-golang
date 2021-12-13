package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func hammingWeightOld(num uint32) int {
	return bits.OnesCount32(num)
}

func hammingWeight(num uint32) int {
	n := strconv.FormatUint(uint64(num), 2)
	fmt.Println(n)

	count := 0
	for _, v := range n {
		if v == '1' {
			count++
		}
	}
	return count
}

func main() {
	var s uint32 = 1100011
	fmt.Println(hammingWeight(s))
}
