package main

import (
	"fmt"
)

func maxChunksToSorted(arr []int) int {
	count, maxLeft := 0, 0
	for i, n := range arr {
		maxLeft = max(maxLeft, n)
		if maxLeft == i {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
