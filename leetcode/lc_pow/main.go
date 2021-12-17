package main

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 1 {
		return x
	} else if x == -1 {
		if n%2 == 0 {
			return -x
		} else {
			return x
		}
	} else if n == 0 {
		return 1
	}

	flag := true
	if n < 0 {
		flag = false
		n *= -1
	}

	res := x
	n--
	for n > 0 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			res *= x
			n--
		}
	}

	if !flag {
		return 1 / res
	}
	return res
}

func main() {
	x := 2.00000
	n := 10
	fmt.Println(myPow(x, n))
}
