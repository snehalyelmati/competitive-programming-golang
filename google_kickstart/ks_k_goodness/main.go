package main

import (
	"fmt"
	"strconv"
)

func calculateK(s string, n int) int {
	count := 0
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			// fmt.Println(i, n-i-1)
			count++
		}
	}

	return count
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var n, k int
		fmt.Scan(&n, &k)

		var s string
		fmt.Scan(&s)

		// calculate k for the string subtract it from desired k
		kString := calculateK(s, n)

		// fmt.Println(kString)
		if k >= kString {
			fmt.Println("Case #"+strconv.Itoa(t+1)+":", k-kString)
		} else {
			fmt.Println("Case #"+strconv.Itoa(t+1)+":", 0)
		}
	}
}
