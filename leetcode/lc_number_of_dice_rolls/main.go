package main

import (
	"fmt"
	"math"
	"strconv"
)

func numRollsToTarget(n int, k int, target int) int {
	return helper("", target, k, n)
}

func helper(p string, u, max, dice int) int {
	fmt.Println(p, u, max, dice)

	if u == 0 {
		fmt.Println("p==target", p)
		return 1
	}

	count := 0
	for i := 1; i <= u && i <= max && dice-1 >= 0; i++ {
		count += helper(p+strconv.Itoa(i), u-i, max, dice-1)
	}

	return count % (int(math.Pow10(9)) + 7)
}

func main() {
	// n, k, target := 1, 6, 3
	n, k, target := 2, 6, 7
	fmt.Println(numRollsToTarget(n, k, target))
}
