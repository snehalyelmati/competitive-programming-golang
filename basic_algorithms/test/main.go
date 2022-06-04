package main

import (
	"fmt"
	"sort"
)

func filterCharacter(s, target string) string {
	if s < target || target == "" {
		return s
	}

	if s[:len(target)] == target {
		return filterCharacter(s[len(target):], target)
	}
	return s[:len(target)] + filterCharacter(s[len(target):], target)
}

func subSequence(p, u string) []string {
	if u == "" {
		fmt.Println(p)
		return []string{p}
	}

	left := subSequence(p+string(u[0]), u[1:])
	right := subSequence(p, u[1:])
	return append(left, right...)
}

func main() {
	// s, target := "applebananapple", ""
	// fmt.Println(filterCharacter(s, target))

	sub := "abc"
	subSeq := subSequence("", sub)
	sort.Strings(subSeq)
	fmt.Println(subSeq)
}
