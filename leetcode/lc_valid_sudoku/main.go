package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for r, row := range board {
		for c, ele := range row {
			// get box coordinates
			rt, ct, rb, cb := getBoxCoordinates(r, c)

			// check row, col and box for all elements where ele != "."
			if string(ele) != "." && (!checkRow(board, r, c, ele) || !checkCol(board, r, c, ele) || !checkBox(board, r, c, rt, ct, rb, cb, ele)) {
				return false
			}
		}
	}
	return true
}

func checkRow(board [][]byte, row, col int, ele byte) bool {
	for i, d := range board[row] {
		if d == ele && i != col {
			fmt.Println("row check failed: ", row, string(ele))
			return false
		}
	}
	return true
}

func checkCol(board [][]byte, row, col int, ele byte) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == ele && i != row {
			fmt.Println("col check failed: ", col, string(ele))
			return false
		}
	}
	return true
}

func checkBox(board [][]byte, r, c, rt, ct, rb, cb int, ele byte) bool {
	for i := rt; i < rb; i++ {
		for j := ct; j < cb; j++ {
			if board[i][j] == ele && i != r && j != c {
				fmt.Println("box check failed: ", string(board[i][j]), i, j, string(ele))
				return false
			}
		}
	}
	return true
}

func getBoxCoordinates(r, c int) (int, int, int, int) {
	// rt, ct, rb, cb
	rt := r - r%3
	ct := c - c%3
	rb := rt + 3
	cb := ct + 3

	return rt, ct, rb, cb
}

func main() {
}
