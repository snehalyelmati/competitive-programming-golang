package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		var even int = 0
		for n := 0; n < N; n++ {
			var i int
			fmt.Scan(&i)

			if i%2==0 {
				even++
			}
		}

		if even > N/2 {
			fmt.Println(N-even)
		} else {
			fmt.Println(even)
		}
	}
}