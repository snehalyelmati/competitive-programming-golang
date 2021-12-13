package main

import "fmt"

func maxProfit(prices []int) int {
	// whole thing should be in a loop for all elements
	profit := 0
	flag := true
	for i := 1; i < len(prices); {
		// fmt.Println(prices[i-1] < prices[i], prices[i-1], prices[i], i)
		if prices[i-1] < prices[i] && flag {
			// buy loop - where it stops descending
			profit -= prices[i-1]
			i++
			// fmt.Println("Profit minus:", profit)
			flag = !flag
		} else if prices[i-1] > prices[i] && !flag {
			// sell loop - where it stops ascending
			profit += prices[i-1]
			i++
			// fmt.Println("Profit plus:", profit)
			flag = !flag
		} else {
			i++
		}
	}
	if !flag {
		profit += prices[len(prices)-1]
	}
	return profit
}

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}

	fmt.Println(prices)
	fmt.Println(maxProfit(prices))
}
