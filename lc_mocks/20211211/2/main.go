package main

import (
	"fmt"
	"sort"
)

func mergeFiles(files []int) int {
	if len(files) < 1 {
		return -1
	}

	sum := 0
	sort.Ints(files)
	for len(files)>1 {
		fmt.Println(files)
		files[1] = files[0]+files[1]
		sum += files[1]
		files = files[1:]
		i := 0
		for i<len(files)-1 && files[i]>files[i+1] {
			files[i], files[i+1] = files[i+1], files[i]
			i++
		}
	}

	return sum
}

func main() {
	files := []int{4, 8, 6, 12}
	fmt.Println(mergeFiles(files))
}




























