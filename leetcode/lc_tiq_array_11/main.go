package main

import "fmt"

func rotate(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := i; j < len(m); j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}

	// printMatrix(m)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m)/2; j++ {
			// comp := len(m) - j - 1
			// fmt.Println(m[i][j], m[i][comp], i, j, ",", i, comp)
			m[i][j], m[i][len(m) - j - 1] = m[i][len(m) - j - 1], m[i][j]
		}
	}

}

func printMatrix(matrix [][]int) {
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println()
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	printMatrix(matrix)
	rotate(matrix)
	printMatrix(matrix)
}
