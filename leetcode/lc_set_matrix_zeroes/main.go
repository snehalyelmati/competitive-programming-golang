package main

import "fmt"

func setZeroes(m [][]int) {
	if len(m) == 0 {
		return
	}
	targetRow := make([]int, 0)
	targetCol := make([]int, 0)
	for j := 0; j < len(m[0]); j++ {
		for i := 0; i < len(m); i++ {
			if m[i][j] == 0 {
				targetRow = append(targetRow, i)
				targetCol = append(targetCol, j)
			}
		}
	}
	for i := 0; i < len(m); i++ {
		for _, v := range targetRow {
			m[v][i] = 0
		}
	}
	for i := 0; i < len(m[0]); i++ {
		for _, v := range targetCol {
			m[i][v] = 0
		}
	}
}

func setZeroesInvalid(m [][]int) {
	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m[0]); i++ {
			if m[j][i] == 0 {
				fmt.Println(j, i)
				fmt.Println(len(m), len(m[0]))
				m[j][0] = 0
				m[0][i] = 0
			}
		}
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for i := 0; i < len(m); i++ {
		if m[i][0] == 0 {
			for j := 0; j < len(m[0]); j++ {
				m[i][j] = 0
			}
		}
	}

	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	for i := 0; i < len(m[0]); i++ {
		if m[0][i] == 0 {
			for j := 0; j < len(m); j++ {
				m[j][i] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 2, 1},
		{3, 0, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)

	for _, v := range matrix {
		fmt.Println(v)
	}
}
