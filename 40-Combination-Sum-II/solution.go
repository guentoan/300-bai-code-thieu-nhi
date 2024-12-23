package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var combination []int

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
		// If target is 0, add the current combination to the result
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		// Try every candidate starting from the current index
		for i := start; i < len(candidates); i++ {
			// Skip duplicates to avoid repeating the same combination
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// Skip if the current number exceeds the target
			if candidates[i] > target {
				break
			}
			// Include the candidate and recurse
			combination = append(combination, candidates[i])
			// Recurse with the next index and reduced target
			backtrack(i+1, target-candidates[i])
			// Backtrack and remove the last element added
			combination = combination[:len(combination)-1]
		}
	}

	// Sort candidates to improve pruning
	sort.Ints(candidates)
	backtrack(0, target)
	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
