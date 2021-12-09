package main

import (
	"fmt"
	"strconv"
)

func removeKdigitsOld(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	minArray := []int{}
	for i := 0; i < len(num); i++ {
		value := int(num[i] - 48)
		if len(minArray) == 0 {
			minArray = append(minArray, value)
		} else if minArray[len(minArray)-1] > value && k > 0 {
			for len(minArray) > 0 && minArray[len(minArray)-1] > value && k > 0 {
				minArray = minArray[:len(minArray)-1]
				k--
			}
			minArray = append(minArray, value)
		} else {
			minArray = append(minArray, value)
		}
	}

	minArray = minArray[:len(minArray)-k]
	for minArray[0] == 0 && len(minArray) > 1 {
		minArray = minArray[1:]
	}
	fmt.Println(minArray)

	res := ""
	for _, d := range minArray {
		res += strconv.Itoa(d)
	}

	return res
}

func removeKdigits(num string, k int) string {
	runes := []rune{}

	for _, d := range num {
		for len(runes) > 0 && d < runes[len(runes)-1] && k > 0 {
			runes = runes[:len(runes)-1]
			k--
		}
		runes = append(runes, d)
	}

	for k > 0 {
		runes = runes[:len(runes)-1]
		k--
	}

	for len(runes) > 0 && runes[0] == '0' {
		runes = runes[1:]
	}

	if len(runes) == 0 {
		return "0"
	}

	return string(runes)
}

func main() {
	// num := "1432219"
	num := "10"
	// num := "00112"
	k := 2
	fmt.Println(removeKdigits(num, k))
}
