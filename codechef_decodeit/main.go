package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		var str string
		fmt.Scan(&str)
		// fmt.Println("String input:", str)

		a := make([]int, N)
		for i := 0; i < N; i++ {
			a[i], _ = strconv.Atoi(string(str[i]))
		}

		for i := 0; i < N; i += 4 {
			value := a[i]*8 + a[i+1]*4 + a[i+2]*2 + a[i+3]*1
			fmt.Print(string(97 + value))
		}
		fmt.Println()
	}
}
