package main

import "fmt" 

func main() {
	var T int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)
		words := make(map[string]struct{})
		firstChars := make(map[string]struct{})
		map1 := make(map[string] []string)
		map2 := make(map[string] []string)
		for i := 0; i < N; i++ {
			var temp string
			fmt.Scan(&temp)
			words[temp] = struct{}{}
			firstChars[temp[0:1]] = struct{}{}
			map1[temp[0:1]] = append(map1[temp[0:1]], temp)
		}
		for k, _ := range words {
			for c, _ := range firstChars {
				temp := c+k[1:]
				_, ok := words[temp]
				if ok == false {
					map2[k] = append(map2[k], c) 
				}
			}
		}
		count := 0
		for k, v := range map2 {
			for _, c := range v {
				for _, word2 := range map1[c] {
					temp := k[0:1] + word2[1:]
					_, ok := words[temp]
					if ok == false {
						count++
					}
				}
			}
		}
		fmt.Println(count)
	}
}