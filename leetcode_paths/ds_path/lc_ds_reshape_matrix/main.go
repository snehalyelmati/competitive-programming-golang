package main

import "fmt"

func matrixReshapeOld(mat [][]int, r int, c int) [][]int {
	ri := len(mat)
	if ri < 1 {
		return mat
	}

	ci := len(mat[0])
	if ri*ci != r*c {
		return mat
	}

	res := make([][]int, r)
	for k := 0; k < r; k++ {
		res[k] = make([]int, c)
	}
	for i := 0; i < ri; i++ {
		for j := 0; j < ci; j++ {
			fmt.Println((i*ri + j), r, c, (i*ri+j)/c, (i*ri+j)%c)
			res[(i*ci+j)/c][(i*ci+j)%c] = mat[i][j]
		}
	}

	return res
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	ri := len(mat)
	if ri < 1 {
		return mat
	}

	ci := len(mat[0])
	if ri*ci != r*c {
		return mat
	}

	res := make([][]int, r)
	for k := 0; k < r; k++ {
		res[k] = make([]int, c)
	}
	k, l := 0, 0
	for i := 0; i < ri; i++ {
		for j := 0; j < ci; j++ {
			res[k][l] = mat[i][j]
			l++
			if l == c {
				k++
				l = 0
			}
		}
	}

	return res
}

func main() {
	mat := [][]int{{1}, {2}, {3}, {4}}
	r, c := 2, 2

	fmt.Println(mat)
	fmt.Println(matrixReshape(mat, r, c))
}
