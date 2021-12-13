package main

import (
	"fmt"
	"strings"
)

func firstUniqCharOld(s string) int {
	hMap := make(map[rune]int)
	for _, v := range s {
		hMap[v] += 1
	}
	fmt.Println(hMap)

	for i, v := range s {
		if hMap[v] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar(s string) int {
	for i, v := range s {
		if strings.Count(s, string(v)) >= 2 {
			continue
		}
		return i
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
