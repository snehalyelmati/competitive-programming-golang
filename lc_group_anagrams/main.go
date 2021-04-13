package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(s1 []string) [][]string {
	res := make([][]string, 0)
	tempStr := ""

	hMap := make(map[string][]string)
	for i, v := range s1 {
		temp := []byte(v)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		// tempStr = strings.Join(temp, "")
		tempStr = string(temp)
		hMap[tempStr] = append(hMap[tempStr], s1[i])
	}
	// fmt.Println(hMap)

	for _, v := range hMap {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
