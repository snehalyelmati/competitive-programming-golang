package main

import "fmt"

func strStr(haystack string, needle string) int {
	ln := len(needle)
	lh := len(haystack)

	if (ln == 0 && lh == 0) || (haystack == needle) {
		return 0
	}

	for i := 0; i < lh-ln+1; i++ {
		fmt.Println(haystack[i:i+ln] == needle[:], haystack[i:i+ln], needle[:], i)
		if haystack[i:i+ln] == needle[:] {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "abc"
	needle := "c"

	fmt.Println(strStr(haystack, needle))
}
