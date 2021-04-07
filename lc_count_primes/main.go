package main

import "fmt"

// primes less than n
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}

	arr := make([]bool, n)
	arr[0] = true
	arr[1] = true

	for i := 2; i < n; i++ {
		if arr[i] == true {
			continue
		}
		for j := 2; i*j < n; j++ {
			arr[i*j] = true
		}
	}

	count := 0
	for _, v := range arr {
		if v != true {
			count++
		}
	}

	return count
}

func main() {
	n := 50
	fmt.Println(countPrimes(n))
}
