package main

import (
	"fmt"
	"slices"
)

func makeStringGood(s string) int {
	// Step 1: Count the frequency of each character in the string
	charFrequency := [26]int{}
	for _, char := range s {
		charFrequency[char-'a']++ // Increment the count for the character
	}

	// Step 2: Find the maximum frequency of any character
	maxFrequency := slices.Max(charFrequency[:]) // This represents the upper bound for the target frequency

	// Step 3: Initialize the minimum cost (answer) to the length of the string
	// This assumes the worst case where we might need to delete the entire string
	minOperations := len(s)

	// Step 4: Prepare a DP array `dp` to store the cost of processing characters
	dp := [27]int{} // DP array, where `dp[i]` is the cost to make all characters from index `i` to 'z' valid

	// Step 5: Iterate over all possible target frequencies
	for targetFrequency := 1; targetFrequency <= maxFrequency; targetFrequency++ {
		// Handle the last character ('z') separately
		dp[25] = min(charFrequency[25], abs(charFrequency[25]-targetFrequency)) // Cost to make 'z' have a frequency of `targetFrequency`

		// Process characters from 'y' to 'a' (backwards)
		for charIndex := 24; charIndex >= 0; charIndex-- {
			currentFrequency := charFrequency[charIndex]
			nextFrequency := charFrequency[charIndex+1]

			// Cost to process the current character to target frequency and add the cost of the next
			dp[charIndex] = dp[charIndex+1] + min(currentFrequency, abs(currentFrequency-targetFrequency))

			// Handle cases where balancing between current and next character frequencies is needed
			if nextFrequency < targetFrequency { // If the next character's frequency is less than target
				if currentFrequency > targetFrequency {
					// Case: Current frequency is higher than target
					// Balance by transferring extra characters from the current group
					dp[charIndex] = min(dp[charIndex], dp[charIndex+2]+max(currentFrequency-targetFrequency, targetFrequency-nextFrequency))
				} else {
					// Case: Current frequency is less than or equal to target
					// Balance by ensuring both groups are valid
					dp[charIndex] = min(dp[charIndex], dp[charIndex+2]+max(currentFrequency, targetFrequency-nextFrequency))
				}
			}
		}

		// Update the overall minimum cost
		minOperations = min(minOperations, dp[0])
	}

	// Return the minimum cost to make the string "good"
	return minOperations
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(makeStringGood("acab"))
	fmt.Println(makeStringGood("wddw"))
	fmt.Println(makeStringGood("aaabbc"))
	fmt.Println(makeStringGood("ruuu"))
}
