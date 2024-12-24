package main

import "fmt"

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func doPartition(s string, start int, current []string, result *[][]string) {
	if start == len(s) {
		p := make([]string, len(current))
		copy(p, current)
		*result = append(*result, p)
		return
	}

	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			current = append(current, s[start:end+1])
			doPartition(s, end+1, current, result)
			// backtracking: remove the last added substring and try the next possibility
			current = current[:len(current)-1]
		}
	}
}

func partition(s string) [][]string {
	var result [][]string
	doPartition(s, 0, []string{}, &result)
	return result
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
