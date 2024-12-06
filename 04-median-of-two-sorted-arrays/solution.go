package main

import (
	"fmt"
	"sort"
)

func solution(nums1, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i := len(nums) / 2
	if len(nums)%2 != 0 {
		return float64(nums[i])
	} else {
		return float64(nums[i-1]+nums[i]) / 2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	res := solution(nums1, nums2)
	fmt.Println(res)
}
