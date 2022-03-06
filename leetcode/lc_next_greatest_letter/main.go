package main

import "fmt"

func nextGreatestLetter(l []byte, target byte) byte {
	low, high := 0, len(l)-1
	idx := (binarySearch(l, low, high, target) + len(l)) % len(l)
	for idx < len(l) && l[idx] == target {
		fmt.Println("Loop:", idx, string(l[idx]), string(target))
		idx++
	}
	return l[(idx+len(l))%len(l)]
}

func binarySearch(l []byte, low, high int, target byte) int {
	mid := (low + high) / 2

	for low <= high {
		if l[mid] > target {
			fmt.Println(string(l[mid]), string(target), l[mid] > target, ">")
			return binarySearch(l, low, mid-1, target)
		} else if l[mid] < target {
			fmt.Println(string(l[mid]), string(target), l[mid] < target, "<")
			return binarySearch(l, mid+1, high, target)
		} else {
			fmt.Println(string(l[mid]), string(target), l[mid] == target, "==")
			return mid
		}
	}

	fmt.Println("Returning low")
	return low
}

func main() {
	letters, target := []byte{'a', 'c', 'f', 'j'}, 'd'
	fmt.Println("Next greatest letter:", string(nextGreatestLetter(letters, byte(target))))
}
