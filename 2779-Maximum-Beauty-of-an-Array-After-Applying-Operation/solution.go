package main

import (
	"fmt"
	"slices"
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	// Sort the array in ascending order
	sort.Ints(nums)

	n := len(nums)
	maxBeauty := 0
	left := 0

	// Use two pointers, left and right, to traverse the array
	for right := 0; right < n; right++ {
		// If the condition is not satisfied, move the left pointer
		for nums[right]-nums[left] > 2*k {
			left++
		}
		// Update the maximum beauty
		maxBeauty = max(maxBeauty, right-left+1)
	}

	return maxBeauty
}

func maximumBeautyHashmap(nums []int, k int) int {
	// Create a hashmap to store the frequency of possible values
	count := make(map[int]int)
	maxBeauty := 0

	for _, num := range nums {
		// For each number, update the range [num-k, num+k] in the hashmap
		for x := num - k; x <= num+k; x++ {
			count[x]++
			// Keep track of the maximum frequency
			if count[x] > maxBeauty {
				maxBeauty = count[x]
			}
		}
	}

	return maxBeauty
}

func maximumBeautyPrefixSum(nums []int, k int) int {
	// Find the maximum value in the array
	mx := slices.Max(nums)

	// Create a difference array with size mx+2
	diff := make([]int, mx+2)

	// Update the difference array based on the range [nums[i]-k, nums[i]+k]
	for _, x := range nums {
		// Mark the start of the range
		diff[max(x-k, 0)]++
		// Mark the end of the range
		diff[min(x+k+1, mx+1)]--
	}

	// Use prefix sum to calculate the count of elements in each range
	count, maxBeauty := 0, 0
	for _, x := range diff {
		count += x
		// Update the maximum beauty of the array
		maxBeauty = max(maxBeauty, count)
	}

	return maxBeauty
}

func main() {
	fmt.Println("Maximum Beauty:", maximumBeauty([]int{4, 6, 1, 2}, 2))  // expect 3
	fmt.Println("Maximum Beauty:", maximumBeauty([]int{1, 1, 1, 1}, 10)) // expect 4

	fmt.Println("Maximum Beauty:", maximumBeautyHashmap([]int{4, 6, 1, 2}, 2))  // expect 3
	fmt.Println("Maximum Beauty:", maximumBeautyHashmap([]int{1, 1, 1, 1}, 10)) // expect 4

	fmt.Println("Maximum Beauty:", maximumBeautyPrefixSum([]int{4, 6, 1, 2}, 2))  // expect 3
	fmt.Println("Maximum Beauty:", maximumBeautyPrefixSum([]int{1, 1, 1, 1}, 10)) // expect 4
}
