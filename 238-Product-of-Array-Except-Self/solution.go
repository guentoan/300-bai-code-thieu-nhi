package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}