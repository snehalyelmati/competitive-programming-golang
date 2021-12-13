package main

import "fmt"

func printBoard(board [][]byte) {
	for _, v := range board {
		fmt.Println(v)
	}
}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		hMap := make(map[byte]bool)
		// fmt.Println(hMap)
		for j := 0; j < 9; j++ {
			// fmt.Println(board[i][j], board[j][i])
			if board[i][j] != '.' {
				// fmt.Println(board[i][j], i, j)
				if _, ok := hMap[board[i][j]]; ok {
					fmt.Println(board[i][j], i, j)
					return false
				} else {
					hMap[board[i][j]] = true
				}
			}
		}
		// fmt.Println(i, hMap)
	}

	// fmt.Println("------------------------------------------")

	for i := 0; i < 9; i++ {
		hMapi := make(map[byte]bool)
		// fmt.Println(hMap)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				// fmt.Println(board[j][i], j, i)
				if _, ok := hMapi[board[j][i]]; ok {
					fmt.Println(board[j][i], j, i)
					return false
				} else {
					hMapi[board[j][i]] = true
				}
			}
		}
		// fmt.Println(i, hMapi)
	}

	fmt.Println("------------------------------------------")

	for k := 0; k < 9; k += 3 {
		hMap1 := make(map[byte]bool)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j+k] != '.' {
					if _, ok := hMap1[board[i][j+k]]; ok {
						return false
					} else {
						hMap1[board[i][j+k]] = true
					}
				}
				// fmt.Print(i, ",", j+k, " ")
			}
			// fmt.Println()
		}
		// fmt.Println()

		hMap2 := make(map[byte]bool)
		for i := 3; i < 6; i++ {
			for j := 0; j < 3; j++ {
				// fmt.Print(i, ",", j+k, " ")
				if board[i][j+k] != '.' {
					if _, ok := hMap2[board[i][j+k]]; ok {
						return false
					} else {
						hMap2[board[i][j+k]] = true
					}
				}
			}
			// fmt.Println()
		}
		// fmt.Println()

		hMap3 := make(map[byte]bool)
		for i := 6; i < 9; i++ {
			for j := 0; j < 3; j++ {
				// fmt.Print(i, ",", j+k, " ")
				if board[i][j+k] != '.' {
					if _, ok := hMap3[board[i][j+k]]; ok {
						return false
					} else {
						hMap3[board[i][j+k]] = true
					}
				}
			}
			// fmt.Println()
		}
		// fmt.Println()
	}

	// for i := 0; i < 9; i += 3 {
	// 	fmt.Println("i:", i)
	// 	hMap := make(map[byte]bool)
	// 	for j := 0; j < 3; j += 1 {
	// 	fmt.Println("j:", j)
	// 		// fmt.Println(i, j)
	// 		fmt.Println(i, j, board[i][j])
	// 		if board[i][j] != '.' {
	// 			// fmt.Println(i, j, board[i][j])
	// 			if _, ok := hMap[board[i][j]]; ok {
	// 				return false
	// 			} else {
	// 				hMap[board[i][j]] = true
	// 			}
	// 		}
	// 		// fmt.Println(i+1, j)
	// 		fmt.Println(i+1, j, board[i+1][j])
	// 		if board[i+1][j] != '.' {
	// 			// fmt.Println(i, j, board[i+1][j])
	// 			if _, ok := hMap[board[i+1][j]]; ok {
	// 				// fmt.Println(board[i][j], i, j)
	// 				return false
	// 			} else {
	// 				hMap[board[i+1][j]] = true
	// 			}
	// 		}
	// 		// fmt.Println(i+2, j)
	// 		fmt.Println(i+2, j, board[i+2][j])
	// 		if board[i+2][j] != '.' {
	// 			// fmt.Println(i, j, board[i+2][j])
	// 			if _, ok := hMap[board[i+2][j]]; ok {
	// 				return false
	// 			} else {
	// 				hMap[board[i+2][j]] = true
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(hMap)
	// }

	return true
}

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '.', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}

	printBoard(board)
	fmt.Println("Is valid sudoku?", isValidSudoku(board))
}
