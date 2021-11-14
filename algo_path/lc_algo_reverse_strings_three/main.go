package main

import (
	"fmt"
	"strings"
)

func reverseString(s []rune) string {
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}

	return string(s)
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")

	for i := 0; i < len(strs); i++ {
		runes := []rune(strs[i])
		strs[i] = reverseString(runes)
		fmt.Println(strs)
	}

	return strings.Join(strs, " ")
}

func reverseWordsOld(s string) string {
	strs := strings.Split(s, " ")
	result := ""

	for i := 0; i < len(strs); i++ {
		runes := []rune(strs[i])
		result += " " + reverseString(runes)
	}

	return result
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
