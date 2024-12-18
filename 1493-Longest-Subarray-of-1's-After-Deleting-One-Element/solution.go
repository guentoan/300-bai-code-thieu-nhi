package main

import "fmt"

func longestSubarray(nums []int) int {
	left, maxLength, zeroCount := 0, 0, 0
	n := len(nums)

	for right := 0; right < n; right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLength = max(maxLength, right-left)
	}

	return maxLength
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}
