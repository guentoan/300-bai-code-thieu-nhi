package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
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
			// Skip if the current number exceeds the target
			if candidates[i] > target {
				break
			}
			// Include the candidate and recurse
			combination = append(combination, candidates[i])
			backtrack(i, target-candidates[i])             // allow repeated numbers
			combination = combination[:len(combination)-1] // backtrack
		}
	}

	// Sort candidates to improve pruning
	sort.Ints(candidates)
	backtrack(0, target)
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
