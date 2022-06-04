package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	for i := 0; i < len(board); i++ {
		helper(&board, 0, i, n, &res)
	}
	return res
}

func helper(board *[][]string, r, c, queens int, res *[][]string) {
	// check if the current space is safe
	// call recursion for the next lines
	// return the board when safetyCheck passes and queens are done

	if !safetyCheck(*board, r, c) {
		// fmt.Println("Failed safetyCheck: ", r, c)
		return
	}
	// fmt.Println("Passed safetyCheck: ", r, c)

	(*board)[r][c] = "Q"
	if queens == 1 {
		*res = append(*res, flattenBoard(*board))
		(*board)[r][c] = "."
		return
	}

	for i := 0; i < len(*board); i++ {
		helper(board, r+1, i, queens-1, res)
	}
	(*board)[r][c] = "."
}

func safetyCheck(board [][]string, r, c int) bool {
	// fmt.Println(board, r, c)
	// check vertical
	for i := 0; i < r; i++ {
		if board[i][c] == "Q" {
			return false
		}
	}

	// check left diagonal
	i, j := r, c
	for i >= 0 && j >= 0 {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}

	// check right diagonal
	i, j = r, c
	for i >= 0 && j < len(board) {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j++
	}

	// all checks passed, return true
	return true
}

func flattenBoard(board [][]string) []string {
	temp := []string{}
	for i := 0; i < len(board); i++ {
		row := ""
		for j := 0; j < len(board); j++ {
			row += board[i][j]
		}
		temp = append(temp, row)
	}
	return temp
}

func printBoard(board [][]string) {
	n := len(board)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
		fmt.Println(board[i])
	}
}

func main() {
	fmt.Println(solveNQueens(5))
}
