package main

import "fmt"

func containsDuplicate(arr []int) bool {
	var hMap = make(map[int]bool, len(arr))
	for _, v := range arr {
		if _, ok := hMap[v]; ok {
			return true
		}
		hMap[v] = true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(arr))
}
