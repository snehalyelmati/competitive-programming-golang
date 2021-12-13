package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var x, y, z int
		fmt.Scanf("%d %d %d", &x, &y, &z)

		if x+y==z || y+z==x || x+z==y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}