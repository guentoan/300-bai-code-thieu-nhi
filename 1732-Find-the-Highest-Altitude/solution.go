package main

import "fmt"

func largestAltitude(gain []int) int {
	high, maxHigh := 0, 0
	n := len(gain)
	for i := 0; i < n; i++ {
		high += gain[i]
		maxHigh = max(maxHigh, high)
	}

	return maxHigh
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
