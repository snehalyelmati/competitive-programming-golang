package main

import "fmt"

func returnSongs(ride int, songs []int) []int {
	songsMap := make(map[int]int)
	for i, d := range songs {
		songsMap[d] = i
	}
	
	for i, d := range songs {
		if idx, ok := songsMap[ride-30-d]; ok && i != idx {
			return []int{i, idx}
		}
	}
	
	return []int{-1, -1}
}

func main() {
	ride := 90
	songs := []int{1, 10, 25, 35, 60}
	
	fmt.Println(returnSongs(ride, songs))
}
