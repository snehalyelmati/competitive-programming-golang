package main

import "fmt"

func numberOfSteps(n int) int {
	return helper(n, 0)
}

func helper(n, count int) int {
	if n == 0 {
		return count
	}

	if n%2 == 0 {
		return helper(n/2, count+1)
	}
	return helper(n-1, count+1)
}

func main() {
	n := 14
	fmt.Println(numberOfSteps(n))
}
