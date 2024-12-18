package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxContainer := 0
	for left < right {
		s := (right - left) * min(height[left], height[right])
		maxContainer = max(maxContainer, s)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxContainer
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
