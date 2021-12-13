package main

import (
	"fmt"
	"regexp"
	"strings"
)

var isAlpNum = regexp.MustCompile("^[a-zA-Z0-9]{1}$").MatchString

func isPalindromeOld(s string) bool {
	res := ""
	for _, v := range s {
		if isAlpNum(strings.ToLower(string(v))) {
			res += strings.ToLower(string(v))
		}
	}

	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		if res[i] != res[len(res)-1-i] {
			return false
		}
	}

	return true
}

func isPalindromeOld2(s string) bool {
	res := ""
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			res += strings.ToLower(string(v))
		}
	}

	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		if res[i] != res[len(res)-1-i] {
			return false
		}
	}

	return true
}

func isAlpNumF(v byte) bool {
	if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	h := 0
	t := len(s) - 1

	for i := 0; i < len(s) && h != t; i++ {
		if isAlpNumF(s[h]) && isAlpNumF(s[t]) {
			fmt.Println(s[h], s[t])
			if strings.ToLower(string(s[h])) != strings.ToLower(string(s[t])) {
				fmt.Println(s[h], s[t])
				return false
			}
			h++
			t--
		} else if isAlpNumF(s[h]) {
			t--
		} else if isAlpNumF(s[t]) {
			h++
		} else {
			h++
			t--
		}
	}
	return true
}

func main() {
	// var s = "A  m man, a plan, a canal: Panama"
	var t = ",; W;:GlG:;l ;,"

	fmt.Println(isPalindrome(t))
}
