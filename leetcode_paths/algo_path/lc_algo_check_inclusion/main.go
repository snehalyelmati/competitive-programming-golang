package main

import (
	"bytes"
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	size1 := len(s1)
	size2 := len(s2)

	if size1 > size2 || size2 == 0 {
		return false
	}
	if size1 == 0 {
		return true
	}

	arr1 := make([]byte, 26)
	arr2 := make([]byte, 26)
	for i := 0; i < size1; i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}

	for i := size1; i < size2; i++ {
		// fmt.Println(string(s2[i]), arr1, arr2)
		if bytes.Equal(arr1, arr2) {
			return true
		}
		arr2[s2[i-size1]-'a']--
		arr2[s2[i]-'a']++
	}

	if bytes.Equal(arr1, arr2) {
		return true
	} else {
		return false
	}
}

func main() {
	s1, s2 := "ab", "eidbiaooo"
	fmt.Println(checkInclusion(s1, s2))
}
