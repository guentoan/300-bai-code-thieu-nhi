package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return count
}

func maxOperations2(nums []int, k int) int {
	hMap := make(map[int]int)
	count := 0

	for _, n := range nums {
		i := k - n
		if hMap[i] > 0 {
			count++
			hMap[i]--
		} else {
			hMap[n]++
		}
	}

	return count
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))

	fmt.Println(maxOperations2([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations2([]int{3, 1, 3, 4, 3}, 6))
}
