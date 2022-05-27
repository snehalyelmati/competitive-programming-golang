package main

import "fmt"

func primeSieve(n int) []int {
	isPrime := make([]int, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = 1
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] == 0 {
			continue
		}
		for j := i; i*j <= n; j++ {
			isPrime[i*j] = 0
		}
	}

	for i := 0; i < len(isPrime); i++ {
		isPrime[i] *= i
	}

	return isPrime
}

func main() {
	n := 100
	fmt.Println(primeSieve(n))
}
