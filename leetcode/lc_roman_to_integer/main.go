package main

import "fmt"

func romanToInt(s string) int {
	sum := 0
	i := 0
	for i = 0; i < len(s)-1; i++ {
		// fmt.Println(string(s[i : i+2]))
		// fmt.Println(string(s[i]))

		if string(s[i:i+2]) == "IV" {
			sum += 4
			i++
			continue
		} else if string(s[i:i+2]) == "IX" {
			sum += 9
			i++
			continue
		} else if string(s[i:i+2]) == "XL" {
			sum += 40
			i++
			continue
		} else if string(s[i:i+2]) == "XC" {
			sum += 90
			i++
			continue
		} else if string(s[i:i+2]) == "CD" {
			sum += 400
			i++
			continue
		} else if string(s[i:i+2]) == "CM" {
			sum += 900
			i++
			continue
		}

		if string(s[i]) == "I" {
			sum += 1
		} else if string(s[i]) == "V" {
			sum += 5
		} else if string(s[i]) == "X" {
			sum += 10
		} else if string(s[i]) == "L" {
			sum += 50
		} else if string(s[i]) == "C" {
			sum += 100
		} else if string(s[i]) == "D" {
			sum += 500
		} else if string(s[i]) == "M" {
			sum += 1000
		}
		// fmt.Println(sum)
	}

	// handle last char case
	if i == len(s)-1 {
		if string(s[i]) == "I" {
			sum += 1
		} else if string(s[i]) == "V" {
			sum += 5
		} else if string(s[i]) == "X" {
			sum += 10
		} else if string(s[i]) == "L" {
			sum += 50
		} else if string(s[i]) == "C" {
			sum += 100
		} else if string(s[i]) == "D" {
			sum += 500
		} else if string(s[i]) == "M" {
			sum += 1000
		}
	}

	return sum
}

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}
