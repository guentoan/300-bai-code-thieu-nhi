package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := 0
	for _, c := range candies {
		if c > maxVal {
			maxVal = c
		}
	}

	res := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= maxVal {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 7}, 3))
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
