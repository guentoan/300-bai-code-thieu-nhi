package main

import "fmt"

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	n := len(s)
	count := 0

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}
	maxCount := count

	for i := k; i < n; i++ {
		if vowels[s[i-k]] {
			count--
		}
		if vowels[s[i]] {
			count++
		}

		maxCount = max(maxCount, count)
	}

	return maxCount
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
}
