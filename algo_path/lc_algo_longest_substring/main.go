package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	hMap := make(map[byte]int)
	max := 0

	for i, j := 0, 0; j < len(s); {
		if idx, ok := hMap[s[j]]; ok && idx >= i {
			i = idx + 1
			continue
		}
		hMap[s[j]] = j
		if max < j-i+1 {
			max = j - i + 1
		}
		j++
	}
	return max
}

func main() {
	// s := "abcabcbb"
	// s := "pwwkew"
	// s := "tmmzuxt"
	s := "abcabcbb"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}
