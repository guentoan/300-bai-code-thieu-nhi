package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for _, n := range nums1 {
		map1[n] = true
	}

	ans := make([][]int, 2)
	for _, n := range nums2 {
		if !map1[n] {
			ans[1] = append(ans[1], n)
			map1[n] = true
		}
		map2[n] = true
	}

	for _, n := range nums1 {
		if !map2[n] {
			ans[0] = append(ans[0], n)
			map2[n] = true
		}
	}

	return ans
}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}
