package main

import "fmt"

/**
* Forward declaration of isBadVersion API.
* @param   version   your guess about first bad version
* @return		      true if current version is bad
*			          false if current version is good
* func isBadVersion(version int) bool;
 */

var badVersion = 1

func isBadVersion(version int) bool {
	return version == badVersion
}

func check(l, h int) int {
	if isBadVersion(l) {
		return l
	} else if l == (l+h)/2 {
		return h
	} else if isBadVersion((l + h) / 2) {
		return check(l, (l+h)/2)
	}
	return check((l+h)/2, h)
}

func firstBadVersion(n int) int {
	return check(1, n)
}

func main() {
	n := 2
	fmt.Printf("Versions: %v, first bad version: %v\n", n, firstBadVersion(n))
}
