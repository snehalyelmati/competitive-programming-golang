package main

import "fmt"

func squareRoot(n int) float64 {
	var mid float64 = float64(n / 2)
	l, h := 0.0, float64(n)

	for l < h {
		mid = (l + h) / 2
		fmt.Println(l, mid, h)
		if mid*mid == float64(n) || mid == l || mid == h {
			return mid
		} else if mid*mid < float64(n) {
			l = mid
		} else if mid*mid > float64(n) {
			h = mid
		}
	}

	return float64(n)
}

func main() {
	n := 400
	fmt.Println(squareRoot(n))
	fmt.Printf("%010.04f\n", squareRoot(n))
}
