package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	helper(&board, 0, 0)
	// fmt.Println(board)
	// fmt.Println(returnStringFromByte(board))
}

func returnStringFromByte(board [][]byte) [][]string {
	res := [][]string{}
	for i := 0; i < len(board); i++ {
		temp := []string{}
		for j := 0; j < len(board); j++ {
			temp = append(temp, string(board[i][j]))
		}
		res = append(res, temp)
	}
	return res
}

func printStringFromBoard(board [][]byte) {
	fmt.Print("[")
	for i := 0; i < len(board); i++ {
		fmt.Print("[")
		for j := 0; j < len(board); j++ {
			fmt.Print("\"" + string(board[i][j]) + "\"")
			if j != len(board)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
		if i != len(board)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func helper(board *[][]byte, r, c int) {
	if c >= 9 {
		r++
		c = 0
	}

	if r == 9 && c == 0 {
		// fmt.Println("In break condition:", r, c)
		printBoard(*board)
		// fmt.Println(returnStringFromByte(*board))
		// printStringFromBoard(*board)
		return
	}

	if string((*board)[r][c]) != "." {
		helper(board, r, c+1)
	} else {
		for i := 1; i <= 9; i++ {
			// fmt.Println(string((*board)[r][c]), string((*board)[r][c]) == ".", safetyCheck(*board, r, c, i), byte(i))
			// fmt.Println(r, c, i)
			if safetyCheck(*board, r, c, i) {
				// fmt.Println("Passed safetyCheck: ", r, c)
				(*board)[r][c] = byte(i + 48)
				// printBoard(*board)
				helper(board, r, c+1)
				(*board)[r][c] = '.'
			}
		}
	}
}

func safetyCheck(board [][]byte, r, c, ch int) bool {
	// vertical check
	for i := 0; i < 9; i++ {
		if board[i][c] == byte(ch+48) {
			return false
		}
	}

	// horizontal check
	for i := 0; i < 9; i++ {
		if board[r][i] == byte(ch+48) {
			return false
		}
	}

	// block check
	row, col := r-r%3, c-c%3
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] == byte(ch+48) {
				return false
			}
		}
	}

	return true
}

func printBoard(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Print(string(board[i][j]))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	printBoard(board)
	fmt.Println()

	solveSudoku(board)
}
