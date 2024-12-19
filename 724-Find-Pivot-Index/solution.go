package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	sumLeft := 0
	for i := 0; i < n; i++ {
		if sumLeft == sum-sumLeft-nums[i] {
			return i
		}

		sumLeft += nums[i]
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
