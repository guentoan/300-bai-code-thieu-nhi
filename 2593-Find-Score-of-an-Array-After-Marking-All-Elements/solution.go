package main

import (
	"fmt"
	"sort"
)

type Element struct {
	value int
	index int
}

func findScore(nums []int) int64 {
	n := len(nums)
	marked := make([]bool, n)
	elements := make([]Element, n)
	var score int64
	for i, num := range nums {
		elements[i] = Element{value: num, index: i}
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].value < elements[j].value || (elements[i].value == elements[j].value && elements[i].index < elements[j].index)
	})

	for _, e := range elements {
		if !marked[e.index] {
			score += int64(e.value)
			marked[e.index] = true

			if e.index > 0 {
				marked[e.index-1] = true
			}
			if e.index < n-1 {
				marked[e.index+1] = true
			}
		}
	}

	return score
}

func main() {
	fmt.Println(findScore([]int{2, 1, 3, 4, 5, 2}))
	fmt.Println(findScore([]int{2, 3, 5, 1, 3, 2}))
}
