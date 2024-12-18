package main

import "fmt"

func finalPrices(prices []int) []int {
	stack := make([]int, 0)
	answer := make([]int, len(prices))
	copy(answer, prices)

	for i, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= price {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[idx] = prices[idx] - price
		}

		stack = append(stack, i)
	}

	return answer
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}
