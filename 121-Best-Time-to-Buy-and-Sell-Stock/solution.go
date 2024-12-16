package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else {
			profit := p - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
