package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	curr := 0
	lower := 0
	hMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if d, ok := hMap[s[i]]; ok && d >= lower {
			lower = d
			curr = i - d
			hMap[s[i]] = i
			fmt.Println(hMap, max, curr, string(s[i]), d, lower)
			continue
		}
		curr++
		if curr > max {
			max = curr
		}
		hMap[s[i]] = i
		fmt.Println(hMap, max, curr, string(s[i]))
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
