package main

import "fmt"

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j < len(t) && i < len(s) {
		if t[j] == s[i] {
			i++
		}
		j++
	}

	return i == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
