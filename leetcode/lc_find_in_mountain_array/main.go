package main

import "fmt"

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, m, h := 0, findPeak(mountainArr), mountainArr.length()-1
	fmt.Println("Peak:", m)
	first := binarySearch(target, mountainArr, l, m)
	if first != -1 {
		return first
	}

	return binarySearch(target, mountainArr, m+1, h)
}

func findPeak(mountainArr *MountainArray) int {
	l, h := 0, mountainArr.length()-1

	for i := 0; l != h; i++ {
		mid := (l + h) / 2

		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			h = mid
		} else {
			l = mid + 1
		}
		fmt.Println(mid)
	}
	return l
}

func binarySearch(target int, mountainArr *MountainArray, l, h int) int {
	flag := 1
	if mountainArr.get(l) > mountainArr.get(h) {
		flag = -1
	}

	for i := l; l < h; i++ {
		mid := (l + h) / 2
		fmt.Println(l, mid, h)
		if flag*mountainArr.get(mid) == flag*target {
			l = mid
			break
		} else if flag*mountainArr.get(mid) > flag*target {
			h = mid - 1
		} else if flag*mountainArr.get(mid) < flag*target {
			l = mid + 1
		}
	}
	if l >= h && flag*mountainArr.get(l) != flag*target {
		return -1
	}
	return l
}

func main() {
}
