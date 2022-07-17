package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	if len(logs) <= 0 {
		return -1
	} else if len(logs) == 1 {
		return logs[0][0]
	}

	// find min
	min := logs[0][0]
	max := logs[0][1]
	for i := 1; i < len(logs); i++ {
		if logs[i][0] < min {
			min = logs[i][0]
		}
		if logs[i][1] > max {
			max = logs[i][1]
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	max -= min
	fmt.Println("Max delta:", max)

	fmt.Println(logs)
	for i := 0; i < len(logs); i++ {
		logs[i][0] -= min
		logs[i][1] -= min
	}
	fmt.Println(logs)

	result := make([]int, max)
	for i := 0; i < len(logs); i++ {
		for j := logs[i][0]; j < logs[i][1]; j++ {
			result[j]++
		}
	}

	count := 0
	year := 0
	for i := 0; i < len(result); i++ {
		if result[i] > count {
			count = result[i]
			year = i
		}
	}

	return min + year
}

func main() {
	// logs := [][]int{{1993, 1999}, {2000, 2010}}
	// logs := [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}
	logs := [][]int{{2008, 2047}, {2018, 2040}, {1953, 1983}}
	fmt.Println(maximumPopulation(logs))
}
