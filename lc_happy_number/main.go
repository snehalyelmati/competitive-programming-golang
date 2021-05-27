package main

import (
	"fmt"
	"math"
)

func isHappyOld(n int) bool {
    hMap := make(map[int]bool)
	for {
		if _, ok := hMap[n]; !ok {
			hMap[n] = true
			n = sumOfSquares(n)
			if n == 1 {
				return true
			}
		} else {
			return false
		}
	}
}

func isHappy(n int) bool {
	slow := sumOfSquares(n)
	fast := sumOfSquares(slow)

	for slow != fast {
		fmt.Println(slow, fast)
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}

	return slow == 1
}

func sumOfSquares(n int) int {
	res := 0
	for n>0 {
		res += int(math.Pow(float64(n%10), float64(2)))
		n /= 10
	}
	return res
}

func main() {
	n := 68
	fmt.Println(isHappy(n))
}