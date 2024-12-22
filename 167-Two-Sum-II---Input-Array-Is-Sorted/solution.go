package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	// Initialize two pointers: left and right
	left, right := 0, len(numbers)-1

	// Loop until the left pointer is less than the right pointer
	for left < right {
		// Calculate the sum of the two numbers at left and right pointers
		sum := numbers[left] + numbers[right]

		// If the sum is equal to the target, return the 1-indexed positions
		if sum == target {
			return []int{left + 1, right + 1} // Add 1 to both indices to match 1-indexed requirement
		}

		// If the sum is less than the target, move the left pointer to the right to increase the sum
		if sum < target {
			left++
		} else { // If the sum is greater than the target, move the right pointer to the left to decrease the sum
			right--
		}
	}

	// Return an empty slice if no solution is found (not expected according to problem statement)
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
