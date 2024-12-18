package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	candies[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}

	return sum
}

func main() {
	fmt.Println(candy([]int{1, 2, 3, 4, 5}))
	fmt.Println(candy([]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 0}))
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}
