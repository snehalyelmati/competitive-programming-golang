package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, M int
		fmt.Scan(&N, &M)
		// fmt.Println(N, M)

		a := make([]int, N)
		aSum := 0
		for i := 0; i < N; i++ {
			fmt.Scan(&a[i])
			aSum += a[i]
		}

		b := make([]int, M)
		bSum := 0
		for i := 0; i < M; i++ {
			fmt.Scan(&b[i])
			bSum += b[i]
		}

		sort.Ints(a)
		sort.Sort(sort.Reverse(sort.IntSlice(b)))

		win := false
		count := 0
		if aSum > bSum {
			win = true
		} else {
			for i := 0; i < int(math.Min(float64(N), float64(M))); i++ {
				a[0], b[0] = b[0], a[0]
				count++

				sort.Ints(a)
				sort.Sort(sort.Reverse(sort.IntSlice(b)))
				
				aSum = 0
				for i := 0; i < N; i++ {
					aSum += a[i]
				}
				
				bSum = 0
				for i := 0; i < M; i++ {
					bSum += b[i]
				}
				
				if aSum > bSum {
					win = true
					break
				}
			}
		}

		if win {
			fmt.Println(count)
		} else {
			fmt.Println(-1)
		}
	}
}
