package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cost(x, y int, s string) int {
	// for all the initial '?' --> fill with the first non '?'
	// for all the ending '?' --> fill with the last non '?'
	// for the next '?' -->
	// 						same before and after --> fill with that
	//						else fill with the cheapest one

	count := 0
	for _, v := range s {
		if v != rune('?') {
			break
		}
		count++
	}
	s = strings.Replace(s, "?", string(s[count]), count)

	for k, v := range s {
		if v == rune('?') && k+1 < len(s) {
			ini := s[:k]
			fin := s[k+1:]
			c := ""
			c = string(s[k+1])
			s = ini + c + fin
		}
	}

	count = 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != "?" {
			break
		}
		count++
	}
	c := string(s[len(s)-1-count])
	// fmt.Println(c, s)
	s = s[:len(s)-count]
	// fmt.Println(s)

	for i := 0; i < count; i++ {
		s += c
	}

	fmt.Println(s)

	return findCost(x, y, s)
}

func findCost(x, y int, s string) int {
	cost := 0
	for i := 0; i < len(s)-1; i++ {
		v := s[i : i+2]
		// fmt.Println(v, string(v))
		if string(v) == "CJ" {
			cost += x
		} else if string(v) == "JC" {
			cost += y
		}
	}
	return cost
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 1; t <= T; t++ {
		var x, y int
		var s string
		fmt.Scan(&x, &y, &s)

		minCost := cost(x, y, s)
		fmt.Println("Case #"+strconv.Itoa(t)+":", minCost)

		// fmt.Println(x, y, s)
	}
}
