package main

import "fmt"

func mySqrtOld(x int) int {
	if x < 2 {
		return x
	}

	res := 0
	for i := 0; i*i <= x; i++ {
		res = i
	}

	return res
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	return helper(x, 0, x/2)
}

func helper(x, l, h int) int {
	mid := (l + h) / 2
	if mid*mid == x || l == h || mid == l {
		return mid
	} else if mid*mid > x {
		return helper(x, l, mid)
	} else {
		return helper(x, mid, h)
	}
}

func main() {
	x := 4

	fmt.Println(mySqrt(x))
}
