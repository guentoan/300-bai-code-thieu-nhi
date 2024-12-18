package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// Initialize the DP table with size (m+1)x(n+1)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: an empty string matches an empty pattern
	dp[0][0] = true

	// Handle the case where the pattern contains '*' at the start
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				// If characters match, or pattern has '.', inherit from dp[i-1][j-1]
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// If pattern contains '*', two cases:
				// 1. Don't use '*' and move 2 steps back in the pattern
				// 2. Use '*' to match one or more characters, check previous row in the table
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	// The final result is stored in dp[m][n], which tells if s matches p
	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
}
