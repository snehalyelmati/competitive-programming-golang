package main

import "fmt"

var res = []string{}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}
	res = []string{}

	curr := ""
	leftRightBal := 0
	leftCount := 0

	helper(n, curr, leftRightBal, leftCount)
	return res
}

func helper(n int, curr string, leftRightBal int, leftCount int) {
	if leftCount == n && leftRightBal == 0 {
		res = append(res, curr)
		return
	}

	if leftCount < n {
		helper(n, curr+"(", leftRightBal+1, leftCount+1)
	}
	if leftRightBal > 0 {
		helper(n, curr+")", leftRightBal-1, leftCount)
	}
}

func main() {
	pNum := 3
	fmt.Println(generateParenthesis(pNum))
}
