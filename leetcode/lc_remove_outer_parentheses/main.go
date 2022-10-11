package main

import "fmt"

func removeOuterParentheses(s string) string {
	count := 0
	for i := 0; i < len(s); {
		if s[i] == '(' {
			if count == 0 {
				s = s[:i] + s[i+1:]
			} else {
				i++
			}
			count++
		} else if s[i] == ')' {
			if count == 1 {
				s = s[:i] + s[i+1:]
			} else {
				i++
			}
			count--
		} else {
			i++
		}
		// fmt.Println(i, s)
	}
	return s
}

func main() {
	s := "(()())(())"
	fmt.Println(removeOuterParentheses(s))
}
