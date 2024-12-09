package main

import (
	"fmt"
	"sort"
)

func isArraySpecial1(nums []int, queries [][]int) []bool {
	type special struct {
		isActive  bool
		isSpecial bool
	}

	points := make([][3]int, 0)
	for i, query := range queries {
		from, to := query[0], query[1]
		points = append(points, [3]int{from, 1, i}) // Start of query
		points = append(points, [3]int{to, 0, i})   // End of query
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	activeQueries := map[int]*special{} // Set of active queries
	pointIdx := 0                       // Pointer for `points`
	for i := 0; i < len(nums)-1; i++ {
		for pointIdx < len(points) && points[pointIdx][0] == i {
			point := points[pointIdx]
			if point[1] == 1 {
				activeQueries[point[2]] = &special{
					isActive:  true,
					isSpecial: true,
				}
			} else {
				activeQueries[point[2]].isActive = false
			}
			pointIdx++
		}

		if nums[i]%2 == nums[i+1]%2 { // Same parity => not special
			for queryIndex := range activeQueries {
				if activeQueries[queryIndex].isActive {
					activeQueries[queryIndex].isSpecial = false
				}
			}
		}
	}

	q := len(queries)
	ans := make([]bool, q)
	for i := range ans {
		if activeQueries[i] != nil {
			ans[i] = activeQueries[i].isSpecial
		} else {
			ans[i] = true
		}

	}

	return ans
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	specials := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		specials[i] = specials[i-1]
		if nums[i-1]%2 != nums[i]%2 {
			specials[i]++
		}
	}

	ans := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		ans[i] = (specials[to] - specials[from]) == (to - from)
	}

	return ans
}

func main() {
	// Example input
	nums := []int{1, 2, 3, 5, 5, 6}
	queries := [][]int{
		{0, 2},
		{1, 4},
		{2, 5},
	}

	// Call the function
	result := isArraySpecial(nums, queries)

	// Print the result
	fmt.Println(result)
}
