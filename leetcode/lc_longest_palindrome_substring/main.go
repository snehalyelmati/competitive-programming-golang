package main

import "fmt"

func longestPalindrome(s string) string {
	maxLen := 0
	// oddFlag := true
	I := 0
	for i := 0; i < len(s); i++ {
		oddLen, oi := expandFromCenter(s, i, i)
		evenLen, ei := expandFromCenter(s, i, i+1)
		// fmt.Println(oi, ei, string(s[i]))

		if oddLen > evenLen && oddLen > maxLen {
			maxLen = oddLen
			// oddFlag = true
			I = oi
		} else if evenLen > oddLen && evenLen > maxLen {
			maxLen = evenLen
			// oddFlag = false
			I = ei
		}
	}

	// return max palindrom string
	return s[I : I+maxLen-2]
}

func expandFromCenter(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l + 1, l + 1
}

func main() {
	s := "baab"
	fmt.Println(longestPalindrome(s))
}
