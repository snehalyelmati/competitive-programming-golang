package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	searchRow := 0
	for _, r := range matrix {
		if r[0] == target {
			return true
		} else if r[0] > target {
			break
		}
		searchRow++ 
	}
	searchRow--
	if searchRow == -1 {
		return false
	}

	for _, d := range matrix[searchRow] {
		if d == target {
			return true
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
