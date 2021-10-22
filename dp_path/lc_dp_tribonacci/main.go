package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	var triboacciArray = make([]int, n+1)
	triboacciArray[0] = 0
	triboacciArray[1] = 1
	triboacciArray[2] = 1

	for i := 3; i <= n; i++ {
		triboacciArray[i] = triboacciArray[i-1] + triboacciArray[i-2] + triboacciArray[i-3]
		fmt.Println(triboacciArray)
	}
	return triboacciArray[n]
}

func main() {
	n := 25
	fmt.Printf("n: %v, tri: %v\n", n, tribonacci(n))
}
