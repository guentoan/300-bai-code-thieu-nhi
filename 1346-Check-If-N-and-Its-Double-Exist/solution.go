package main

import "fmt"

func checkIfExist(arr []int) bool {
	// Create a map to store the numbers we've encountered in the array.
	exists := make(map[int]bool)

	// Iterate over each element in the array.
	for _, v := range arr {
		// Check if the double of the current value (v*2) is in the map,
		// or if the value is even and half of the value (v/2) exists in the map.
		// If either condition is true, return true.
		if exists[v*2] || (v%2 == 0 && exists[v/2]) {
			return true
		}
		// If no such pair is found, add the current value to the map.
		exists[v] = true
	}

	// If no pair is found after checking all elements, return false.
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{3, 1, 7, 11})) // expected false
	fmt.Println(checkIfExist([]int{33, 1, 2, 3})) // expected true
}
