package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	hMap := make(map[string]bool)
	return check(s, wordDict, hMap)
}

func check(s string, wordDict []string, hMap map[string]bool) bool {
	if s == "" {
		return true
	}

	if data, ok := hMap[s]; ok {
		return data
	}

	if wordInDict(s, wordDict) {
		hMap[s] = true
		return true
	}

	for i := 1; i < len(s); i++ {
		if wordInDict(s[:i], wordDict) && check(s[i:], wordDict, hMap) {
			hMap[s[i:]] = true
			return true
		}
	}

	hMap[s] = false
	return false
}

// not implemented
func wordBreakOld(s string, wordDict []string) bool {
	size := len(s)
	valid := make([][]bool, size)
	for i := range valid {
		valid[i] = make([]bool, size)
	}

	fmt.Println(s, size, wordDict)

	for l := 0; l <= size; l++ {

		// whole word loop
		for i := 0; i < size-l; i++ {
			if wordInDict(s[i:i+l+1], wordDict) {
				fmt.Println(s[i : i+l+1])
				valid[i][i+l] = true
				continue
			}

			// substring of length l loop
			for j := i + 1; j < i+l; j++ {
				// check if word(i:i+l+1) is true
				// check if word(i:j) and word(j:i+l) are true
				// otherwise check if valid[i][j] and valid[j][i+l] are true
				// if condition one or two is true for both cases then set valid[i][i+l] true

				// fmt.Printf("Full string now: %v, (%v, %v), %v\n", s[i:i+l+1], i, i+l+1, l+1)
				// fmt.Println(s[i:j], wordInDict(s[i:j], wordDict), valid[i][j])
				// fmt.Println(s[j:i+l], wordInDict(s[j:i+l], wordDict), valid[j][i+l])

				// } else if (wordInDict(s[i:j], wordDict) || valid[i][j]) && (wordInDict(s[j:i+l+1], wordDict) || valid[j][i+l]) {
				if valid[i][j] && valid[j][i+l] {
					fmt.Println(s[i:j+1], s[j:i+l+1])
					// fmt.Println((wordInDict(s[i:j], wordDict) || valid[i][j]) && (wordInDict(s[j:i+l+1], wordDict) || valid[j][i+l]))
					valid[i][i+l] = true
				}
				// fmt.Println()
			}
		}
	}

	// print the validity matrix
	fmt.Println("Validity matrix:")
	for i := range valid {
		fmt.Println(valid[i])
	}

	// if valid[0][size-1] is true return true
	return valid[0][size-1]
}

func wordInDict(word string, list []string) bool {
	for _, ele := range list {
		if ele == word {
			return true
		}
	}
	return false
}

func main() {
	str1 := "leetcode"
	wordDict1 := []string{"leet", "code"}

	str2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}

	str3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}

	fmt.Println(wordBreak(str1, wordDict1))
	fmt.Println(wordBreak(str2, wordDict2))
	fmt.Println(wordBreak(str3, wordDict3))
}
