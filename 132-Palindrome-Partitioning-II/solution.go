package main

import "fmt"

func minCut(s string) int {
	n := len(s)

	// cuts[i] will store the minimum cuts required to partition s[0..i] into palindromes
	cuts := make([]int, n)

	// Initialize cuts array with worst case: cuts[i] = i (because worst case we can cut after each character)
	for i := 0; i < n; i++ {
		cuts[i] = i
	}

	// A helper 2D array to check if s[i..j] is a palindrome
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Check for all substrings and update the cuts array
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || isPalindrome[j+1][i-1]) {
				isPalindrome[j][i] = true
				if j == 0 {
					cuts[i] = 0
				} else {
					cuts[i] = min(cuts[i], cuts[j-1]+1)
				}
			}
		}
	}

	return cuts[n-1]
}

func main() {
	fmt.Println(minCut("babad"))
	fmt.Println(minCut("aab"))
	fmt.Println(minCut("a"))
	fmt.Println(minCut("ab"))
}
