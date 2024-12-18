package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < n; i++ {
		sum = sum + nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
