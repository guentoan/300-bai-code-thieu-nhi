package main

import "fmt"

func maximumLength(s string) int {
	// Step 1: Initialize map to store counts
	countMap := make(map[string]int)

	// Step 2: Traverse the string to find special substrings
	n := len(s)
	i := 0
	for i < n {
		// Find a sequence of identical characters
		start := i
		for i < n && s[i] == s[start] {
			i++
		}

		// Calculate all special substrings within this sequence
		length := i - start // Length of the special sequence
		for l := 1; l <= length; l++ {
			substr := s[start : start+l]
			countMap[substr] += length - l + 1
		}
	}

	// Step 3: Find the longest special substring occurring at least 3 times
	maxLength := -1
	for substr, count := range countMap {
		if count >= 3 && len(substr) > maxLength {
			maxLength = len(substr)
		}
	}

	return maxLength
}

func main() {
	// Example test cases
	fmt.Println(maximumLength("aaaabbb"))   // Output: 2 ("aa")
	fmt.Println(maximumLength("aabb"))      // Output: -1 (no substring occurs at least 3 times)
	fmt.Println(maximumLength("zzzzzzzz"))  // Output: 6 ("zzzzzz")
	fmt.Println(maximumLength("abcabcabc")) // Output: 1 ("a")
}
