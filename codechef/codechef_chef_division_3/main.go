package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, K, D int
		fmt.Scanf("%d %d %d", &N, &K, &D)

		a := make([]int, N)
		aSum := 0
		for i := 0; i < N; i++ {
			fmt.Scan(&a[i])
			aSum += a[i]
		}

		maxDays := aSum / K

		if maxDays > D {
			fmt.Println(D)
		} else {
			fmt.Println(maxDays)
		}
	}
}
