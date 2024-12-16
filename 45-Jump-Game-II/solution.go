package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	maxReach, current, count := 0, 0, 0

	for i := 0; i < n-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == current {
			count++
			current = maxReach
		}
	}

	return count
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}
