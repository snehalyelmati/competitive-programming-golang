package main

import (
	"fmt"
	"math"
	"strconv"
)

func maximalSquare(m [][]byte) int {
	res := make([][]int, len(m)+1)
	for i := range res {
		res[i] = make([]int, len(m[0])+1)
	}

	maxSide := 0
	for i:=0; i<len(m); i++ {
		for j:=0; j<len(m[0]); j++ {
			if data, _ := strconv.Atoi(string(m[i][j])); data == 0 {
				continue
			}
			printMatrix(res)

			lt := res[i][j]
			rt := res[i][j+1]
			lb := res[i+1][j]

			res[i+1][j+1] = findMin([]int{lt, rt, lb})+1

			if maxSide < res[i+1][j+1] {
				maxSide = res[i+1][j+1]
			}
		}
	}

	fmt.Println("Maximal square end...")
	printMatrix(res)
	return maxSide*maxSide
}

func findMin(a []int) int {
	fmt.Println("findMin:", a)
	min := math.MaxInt32
	for _, ele := range a {
		if ele < min {
			min = ele
		}
	}
	return min
}

func printMatrix(m [][]int) {
	for _, row := range m {
		for _, ele := range row {
			fmt.Print(ele, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMatrixByte(m [][]byte) {
	for _, row := range m {
		for _, ele := range row {
			fmt.Print(string(ele), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	matrix := [][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	}
	printMatrixByte(matrix)

	fmt.Println(maximalSquare(matrix))
}

