package main

import "fmt"

func topKFrequent(n []int, k int) []int {
	hMap := make(map[int]int)
	for _, v := range n {
		hMap[v]++
	}

	res := make([]int, 0)
	maxEle := 0
	maxCount := 0
	for i := 0; i < k; i++ {
		maxEle = 0
		maxCount = 0
		for k, v := range hMap {
			if v > maxCount {
				maxEle = k
				maxCount = v
			}
		}

		res = append(res, maxEle)
		hMap[maxEle] = 0
	}

	return res
}

func main() {
	n := []int {1,1,1,2,2,3,3,3,3}
	k := 2

	fmt.Println(topKFrequent(n, k))
}