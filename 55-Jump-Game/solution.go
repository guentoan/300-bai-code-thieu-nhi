package main

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0
	for i, n := range nums {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+n)
	}

	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
