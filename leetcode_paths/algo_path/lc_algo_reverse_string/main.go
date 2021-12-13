package main

import "fmt"

func reverseString(s []byte) {
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
}

func main() {
	s := []byte{'w', 'o', 'h', 'e', 'l', 'l', 'o'}

	fmt.Println(string(s))
	reverseString(s)
	fmt.Println(string(s))
}
