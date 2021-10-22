package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	n := 6
	fmt.Printf("n: %v, fib: %v\n", n, fib(n))
}
