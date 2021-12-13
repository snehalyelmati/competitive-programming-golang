package main

import "fmt"

func generate(rows int) [][]int {
	arr := make([][]int, 0)
	arr = append(arr, []int{1})
	if rows == 1 {
		return arr
	}
	arr = append(arr, []int{1, 1})
	if rows == 2 {
		return arr
	}

	for i := 2; i < rows; i++ {
		temp := make([]int, 0)
		temp = append(temp, 1)
		for j := 0; j < len(arr[i-1])-1; j++ {
			temp = append(temp, arr[i-1][j]+arr[i-1][j+1])
		}
		temp = append(temp, 1)
		arr = append(arr, temp)
	}

	return arr
}

func main() {
	rows := 5

	fmt.Println(generate(rows))
}
