package main

import "fmt"

func setZeroesOptimal(matrix [][]int)  {
    if len(matrix) == 0 {
        return
    }
    fmt.Println("Original matrix:")
    printMatrix(matrix)
    
    rowZ, colZ := false, false
    for i:=0; i<len(matrix); i++ {
        for j:=0; j<len(matrix[0]); j++ {
            if i == 0 || j == 0 {
                if matrix[i][j] == 0 && i == 0 {
                    rowZ = true
                }
                if matrix[i][j] == 0 && j == 0 {
                    colZ = true
                }
                continue
            }
            
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }
    fmt.Println("Mod matrix:")
    printMatrix(matrix)
    
    for i:=1; i<len(matrix[0]); i++ {
        if matrix[0][i] == 0 {
            setCol(matrix, i)
        }
    }
    fmt.Println("Col matrix:")
    printMatrix(matrix)
    
    for i:=1; i<len(matrix); i++ {
        if matrix[i][0] == 0 {
            setRow(matrix, i)
        }
    }
    fmt.Println("Row matrix:")
    printMatrix(matrix)
    
    if rowZ == true {
        setRow(matrix, 0)
    }
    if colZ == true {
        setCol(matrix, 0)
    }
    fmt.Println("Result matrix:")
    printMatrix(matrix)
}

func setRow(matrix [][]int, r int) {
    for i:=0; i<len(matrix[0]); i++ {
        matrix[r][i] = 0
    }
}

func setCol(matrix [][]int, c int) {
    for i:=0; i<len(matrix); i++ {
        matrix[i][c] = 0
    }
}

func printMatrix(mat [][]int) {
    for _, di := range mat {
        for _, dj := range di {
            fmt.Print(dj, " ")
        }
        fmt.Println()
    }
}

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
