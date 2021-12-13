package main

import "fmt"

func titleToNumber(columnTitle string) int {
	res := 0
	for _, v := range columnTitle {
		fmt.Println(int(v), int(v-64))
		res = 26*res + int(v-64)
	}
	return res
}

func main() {
	es := "AY"
	fmt.Println(titleToNumber(es))
}
