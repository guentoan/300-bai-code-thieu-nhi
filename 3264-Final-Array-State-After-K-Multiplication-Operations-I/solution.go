package main

import "fmt"

func getFinalState(nums []int, k int, multiplier int) []int {
	n := len(nums)
	for i := 0; i < k; i++ {
		val := nums[0]
		idx := 0
		for j := 1; j < n; j++ {
			if nums[j] < val {
				val = nums[j]
				idx = j
			}
		}
		nums[idx] = val * multiplier
	}

	return nums
}

func main() {
	fmt.Println(getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Println(getFinalState([]int{1, 2}, 3, 4))
}
