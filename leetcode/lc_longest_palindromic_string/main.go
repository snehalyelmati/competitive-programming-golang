package main

import "fmt"

func longestPalindrome(s string) string {
	maxPalindromeSize := 0
	palindromeCenter := 0
	// nextBoundary := 0
	palindrome := ""

	for i := 0; i < len(s); i++ {
		diff := 1
		palindromeSize := 1
		l := i
		r := i + 1
		for l-diff >= 0 && r+diff < len(s) && s[l] == s[r] {
			fmt.Println(string(s[l]), string(s[r]))
			if s[l-diff] == s[r+diff] {
				diff++
				palindromeSize += 1
				// nextBoundary = i + diff
				if palindromeSize > maxPalindromeSize {
					maxPalindromeSize = palindromeSize
					palindromeCenter = i
				}
				continue
			}
			// i = nextBoundary
			// diff++
			break
		}
	}
	// maxPalindromeSize
	fmt.Println(palindromeCenter, maxPalindromeSize)
	palindrome = s[palindromeCenter-maxPalindromeSize : palindromeCenter+1+maxPalindromeSize]

	maxPalindromeSize = 1
	palindromeCenter = 0
	for i := 0; i < len(s); i++ {
		diff := 1
		palindromeSize := 1
		for i-diff >= 0 && i+diff < len(s) {
			if s[i-diff] == s[i+diff] {
				diff++
				palindromeSize += 1
				// nextBoundary = i + diff
				if palindromeSize > maxPalindromeSize {
					maxPalindromeSize = palindromeSize
					palindromeCenter = i
				}
				continue
			}
			// i = nextBoundary
			// diff++
			break
		}
	}

	maxPalindromeSize--
	fmt.Println(palindromeCenter, maxPalindromeSize)
	palindrome2 := s[palindromeCenter-maxPalindromeSize : palindromeCenter+maxPalindromeSize+1]

	fmt.Println(palindrome, palindrome2)
	if len(palindrome) > len(palindrome2) {
		return palindrome
	}
	return palindrome2
}

func main() {
	s := "caad"
	fmt.Println(longestPalindrome(s))
}
