package main

import (
	"bufio"
	"fmt"
	"os"
)

// Abs returns the absolute value of x.
// func Abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// Min returns minimum value.
// func Min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}
func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func main() {
	defer writer.Flush()
	
	var T int
	scanf("%d", &T)
	
	var N, K, x, y int
	for t := 0; t < T; t++ {
		scanf("%d %d %d %d\n", &N, &K, &x, &y)

		if x == y {
			printf("%d %d\n", N, N)
			continue
		}

		cond := y < x

		// diff := Min(N-x, N-y)

		if cond {
			diff := N - x
			x += diff
			y += diff
		} else {
			diff := N - y
			x += diff
			y += diff
		}

		if K%4 == 1 {
			printf("%d %d\n", x, y)
			continue
		} else if K%4 == 2 {
			printf("%d %d\n", y, x)
			continue
		} else if K%4 == 3 && cond {
			printf("%d %d\n", 0, x-y)
			continue
		} else if K%4 == 3 {
			printf("%d %d\n", y-x, 0)
			continue
		} else if cond {
			printf("%d %d\n", x-y, 0)
			continue
		} else {
			printf("%d %d\n", 0, y-x)
			continue
		}
	}
}
