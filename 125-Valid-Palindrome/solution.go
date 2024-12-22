package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	// Initialize two pointers: left and right
	left, right := 0, len(s)-1

	// Loop until the two pointers meet
	for left < right {
		// Move left pointer to the next alphanumeric character
		if !isAlphanumeric(rune(s[left])) {
			left++
			continue
		}

		// Move right pointer to the previous alphanumeric character
		if !isAlphanumeric(rune(s[right])) {
			right--
			continue
		}

		// Convert both characters to lowercase and compare them
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false // If characters do not match, return false
		}

		// Move both pointers inward
		left++
		right--
	}

	// If we have checked all characters, return true
	return true
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}
