package main

import "fmt"

func plusOne(arr []int) []int {
	flag := false
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < 9 {
			arr[i]++
			flag = true
			break
		} else {
			arr[i] = 0
		}
	}
	if !flag {
		arr = append([]int{1}, arr...)
	}
	return arr
}

func main() {
	arr := []int{9}
	fmt.Println(plusOne(arr))
}
