package main

import "fmt"

var result = make([]string, 0)
func letterCombinations(digits string) []string {
	result = []string{}

	if len(digits) == 0 {
		return result
	}

	mapping := []string{
		"0",
		"1",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	letterCombinationsRecursive(digits, "", 0, mapping)
	return result
}

func letterCombinationsRecursive(digits string, curr string, id int, mapping []string) {
	if id == len(digits) {
		result = append(result, curr)
		return
	}

	letters := mapping[digits[id]-'0']
	for _, v := range letters {
		letterCombinationsRecursive(digits, curr + string(v), id+1, mapping)
	}
}

func main() {
	digits := ""
	fmt.Println(letterCombinations(digits))
}
