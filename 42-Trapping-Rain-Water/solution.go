package main

import "fmt"

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight, count := 0, 0, 0

	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				count += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				count += maxRight - height[right]
			}
			right--
		}
	}

	return count
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
