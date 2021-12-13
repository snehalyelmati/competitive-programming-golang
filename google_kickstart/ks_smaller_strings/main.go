package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	limit := int(math.Pow(10.0, 9.0)) + 7

	for t := 0; t < T; t++ {
		var N, K int
		fmt.Scanf("%d %d", &N, &K)

		var str, oriStr string
		fmt.Scan(&str)
		oriStr = str
		// fmt.Println(str, N, K)

		mid := 0
		if N%2 != 0 {
			mid = N/2 + 1
		} else {
			mid = N / 2
		}

		maxComb := int(math.Pow(float64(K), float64(mid)))
		for i := 0; i <= N/2; i++ {
			comp := N - i - 1
			if str[i] < str[comp] {
				str = str[:comp] + string(str[i]) + str[comp+1:]
			} else if str[i] > str[comp] {
				str = str[:i] + string(str[comp]) + str[i+1:]
			}
		}

		if str == oriStr && len(str) == 1 {
			str = string(str[0] - 1)
		} else if str == oriStr {
			str = str[:mid] + string(str[mid]-1) + str[mid+1:]
		}

		// fmt.Println("Max palindrom:", str)
		// fmt.Println("Max combo and mid:", maxComb, mid)
		// fmt.Println("str[i]:", N-int(str[0]-96))

		res := maxComb
		for i := 0; i <= N/2; i++ {
			temp := (K - int(str[i]-96)) * int(math.Pow(float64(K), float64(mid-i-1)))
			// fmt.Println(tempA, tempB, tempA*tempB)
			res -= temp
			res %= limit
		}

		fmt.Printf("Case #%d: %v\n", t+1, res)
	}
}

// 5*5*5 = 125
// 4*5*5 + 3*5 + 2 = 100 + 15 + 2 = 117
