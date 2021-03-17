package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	} else if n == 3 {
		return "21"
	} else if n == 4 {
		return "1211"
	}
	res := "1211"

	for i := 4; i < n; i++ {
		res = makeRes(res)
		// fmt.Println("RES:", res)
	}

	return res
}

func makeRes(s string) string {
	flag := false
	res := ""
	count := 1
	i := 0
	for i = 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
			flag = false
			// fmt.Println(count)
		} else {
			// fmt.Println(strconv.Itoa(count), string(s[i]))
			res += strconv.Itoa(count) + string(s[i])
			count = 1
			flag = true

			if i == len(s)-2 {
				res += "1" + string(s[len(s)-1])
			}
		}
	}
	// if i == len(s)-1 {
	// 	res += "1" + string(s[i])
	// 	flag = true
	// }
	if !flag {
		res += strconv.Itoa(count) + string(s[i])
	}
	return res

	// hmap := make(map[string]int)
	// for _, v := range s {
	// 	hmap[string(v)]++
	// }

	// fmt.Println(hmap, s)

	// s = removeDuplicates(s)
	// return s

	// res := ""
	// for _, v := range s {
	// 	fmt.Println("Adding this:", strconv.Itoa(hmap[string(v)])+string(v), string(v))
	// 	res += strconv.Itoa(hmap[string(v)]) + string(v)
	// }

	// return res
}

func removeDuplicates(s string) string {
	count := 1
	res := string(s[0]) + ""
	for _, v := range s {
		if string(v) != string(res[len(res)-1]) {
			res += strconv.Itoa(count) + string(v)
		} else {
			count++
		}
	}
	return res
}

func main() {
	n := 10

	fmt.Println(countAndSay(n))
}
