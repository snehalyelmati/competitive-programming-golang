package main

import "fmt"

func trailingZeroes(n int) int {
	i := 5
	count := 0
	for n/i >= 1 {
		count += n/i
		i *= 5
	}
	return count
}

func main() {
	x := 50
	fmt.Println(trailingZeroes(x))
}