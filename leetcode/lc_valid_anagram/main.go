package main

import "fmt"

func isAnagram(s string, t string) bool {
	hMap := make(map[rune]int)
	for _, v := range s {
		hMap[v]++
	}

	for _, v := range t {
		if _, ok := hMap[v]; !ok {
			return false
		}
		if hMap[v] == 1 {
			delete(hMap, v)
		} else {
			hMap[v]--
		}
	}

	return len(hMap) == 0
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
