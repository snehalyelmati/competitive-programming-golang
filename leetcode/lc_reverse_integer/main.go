package main

import (
	"fmt"
	"strconv"
)

var max = 2147483647
var min = -2147483648

func reverse(x int) int {
	s := strconv.Itoa(x)

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	fmt.Println(res)
	if string(res[len(s)-1]) == "-" {
		res = string(res[len(s)-1]) + res[:len(s)-1]
	}
	// fmt.Println(res)

	x, _ = strconv.Atoi(res)

	if x < min || x > max {
		return 0
	}

	return x
}

func main() {
	x := 120

	fmt.Println(reverse(x))
}
