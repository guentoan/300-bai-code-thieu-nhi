package main

import (
	"fmt"
	"sort"
)

type pair struct{ height, index int }

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	newQueries := make([][]pair, n)
	ans := make([]int, len(queries))

	// Preprocess queries
	for i, query := range queries {
		ai, bi := query[0], query[1] // Extract the indices of Alice and Bob's buildings
		// Ensure ai < bi for consistent processing
		if ai > bi {
			ai, bi = bi, ai
		}

		// If Alice and Bob are at the same building or Alice's building is shorter
		if ai == bi || heights[ai] < heights[bi] {
			ans[i] = bi // Alice and Bob can meet at Bob's building
		} else {
			// If Alice's building is taller, store this query for later processing
			newQueries[bi] = append(newQueries[bi], pair{heights[ai], i})
		}
	}

	// Stack to keep track of buildings in decreasing order of height (right to left)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		// Process queries related to Bob's building at position i
		for _, q := range newQueries[i] {
			// Find the nearest building in the stack taller than Alice's building
			j := sort.Search(len(stack), func(i int) bool { return heights[stack[i]] <= q.height }) - 1
			if j == -1 {
				// If no suitable building is found, the result is -1
				ans[q.index] = -1
			} else {
				// Assign the result to the nearest suitable building
				ans[q.index] = stack[j]
			}
		}

		// Remove buildings from the stack that are shorter than or equal to the current building
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
		}

		// Add the current building to the stack
		stack = append(stack, i)
	}

	return ans
}

func main() {
	fmt.Println(leftmostBuildingQueries([]int{6, 4, 8, 5, 2, 7}, [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}))       //[2,5,-1,5,2]
	fmt.Println(leftmostBuildingQueries([]int{5, 3, 8, 2, 6, 1, 4, 6}, [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}})) //[7,6,-1,4,6]
}
