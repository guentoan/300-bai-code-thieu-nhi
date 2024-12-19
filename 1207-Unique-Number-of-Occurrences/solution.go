package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, n := range arr {
		occurrences[n]++
	}

	seen := make(map[int]struct{})
	for _, count := range occurrences {
		if _, existed := seen[count]; existed {
			return false
		}
		seen[count] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
