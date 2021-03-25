package main

import (
	"fmt"
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersionOld(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return -1
}

func check(f, l int) int {
	fmt.Println(f, l)
	if isBadVersion(f) {
		// fmt.Println("Returning here...")
		return f
	} else if f == f+(l-f)/2 {
		return l
	} else if isBadVersion(f + (l-f)/2) {
		// fmt.Println("Returning here...")
		return check(f, f+(l-f)/2)
	}

	return check(f+(l-f)/2, l)
}

func firstBadVersion(n int) int {
	// Binary search
	// check first -> check mid (don't check last)
	// if mid != bad -> check right half of the array

	return check(1, n)
}

var b int = 1702766719

func isBadVersion(n int) bool {
	return n >= b
}

func main() {
	n := 2126753390
	fmt.Println(firstBadVersion(n))
}
