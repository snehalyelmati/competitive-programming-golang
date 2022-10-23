package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	// records traversal
	boolMatrix := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		boolMatrix[i] = make([]bool, len(matrix[0]))
	}
	printMatrix(boolMatrix, "Bool matrix")

	i, j := 0, 0
	rowFlag, colFlag := 0, 0
	sw := false

	for n := 0; n < len(matrix)*len(matrix[0]); n++ {
		// append current element
		res = append(res, matrix[i][j])

		// mark the current element
		boolMatrix[i][j] = true

		// compute next element and increment
		fmt.Println(res, rowFlag, colFlag, i, j)

		// check if existing path is valid
		if (i+rowFlag >= 0 && i+rowFlag < len(matrix) && !boolMatrix[i+rowFlag][j]) ||
			(j+colFlag >= 0 && j+colFlag < len(matrix[0]) && !boolMatrix[i][j+colFlag]) {
			if sw {
				i += rowFlag
			} else {
				j += colFlag
			}
			continue
		}

		// find new path
		fmt.Println("Finding new path:")
		if j+1 < len(matrix[0]) && !boolMatrix[i][j+1] {
			// check right
			j++
			colFlag = 1
			sw = false
		} else if i+1 < len(matrix) && !boolMatrix[i+1][j] {
			// check down
			i++
			rowFlag = 1
			sw = true
		} else if j-1 >= 0 && !boolMatrix[i][j-1] {
			// check left
			j--
			colFlag = -1
			sw = false
		} else if i-1 >= 0 && !boolMatrix[i-1][j] {
			// check up
			i--
			rowFlag = -1
			sw = true
		} else {
			// if everything is covered => break
			break
		}
	}

	return res
}

func printMatrix[T int | bool](mat [][]T, name string) {
	fmt.Println(name + ":")
	for _, di := range mat {
		for _, dj := range di {
			fmt.Print(dj, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	printMatrix(matrix, "Original")
	fmt.Println(spiralOrder(matrix))
}
