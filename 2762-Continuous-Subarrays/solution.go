package main

import "fmt"

func continuousSubarrays(nums []int) int64 {
	start := 0
	n := len(nums)
	window := []int{}
	var result int64

	for end := 0; end < n; end++ {
		window = append(window, nums[end])
		for maxArr(window)-minArr(window) > 2 {
			window = window[1:]
			start++
		}

		result += int64(end - start + 1)
	}

	return result
}

func maxArr(nums []int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func minArr(nums []int) int {
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func main() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4}))
	fmt.Println(continuousSubarrays([]int{1, 2, 3}))
}
