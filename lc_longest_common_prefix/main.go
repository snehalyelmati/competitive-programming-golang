package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	fmt.Println(strs)
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println(strs)

	count := 0
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	for i := len(strs[0]); i > 0; i-- {
		for _, v := range strs {
			// fmt.Println(strs[0][:i], v[:i])
			if strs[0][:i] == v[:i] {
				count++
				continue
			} else {
				break
			}
		}
		if count == len(strs) {
			return strs[0][:i]
		}
		count = 0
	}

	return ""
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
