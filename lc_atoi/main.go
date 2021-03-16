package main

import "fmt"

var max = 2147483647
var min = -2147483648

func myAtoiOld(s string) int {
	var flag, sign bool
	var res int
	for _, v := range s {
		if sign && (v == '-' || v == '+' || v == ' ') {
			break
		}

		if v == ' ' {
			continue
		}

		if v == '+' {
			flag = false
			sign = true
			continue
		} else if v == '-' {
			flag = true
			sign = true
			continue
		}

		if v < '0' || v > '9' {
			break
		}

		sign = true
		res = res*10 + int(v) - 48
	}
	if flag {
		res *= -1
	}

	if res < min {
		res = min
	} else if res > max {
		res = max
	}

	return res
}

func myAtoi(s string) int {
	i := 0
	res := 0

	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else {
			break
		}
	}

	if i == len(s) {
		return res
	}

	sign := false
	if s[i] == '-' {
		sign = true
		i++
	} else if s[i] == '+' {
		i++
	}

	fmt.Println(res)

	for j := i; j < len(s); j++ {
		fmt.Println(res)
		if s[j] >= '0' && s[j] <= '9' {
			res = res*10 + int(s[j]) - 48
			if sign && -res < min {
				fmt.Println(res)
				return min
			} else if !sign && res > max {
				return max
			}
		} else {
			break
		}
	}

	if sign {
		return -res
	}
	return res
}

func main() {
	s := "000123"
	fmt.Println(myAtoi(s))
}
