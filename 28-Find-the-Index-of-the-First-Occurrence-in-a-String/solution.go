package main

import "fmt"

func strStr(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	if n > h {
		return -1
	}

	for i := 0; i <= h-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}

func buildLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func strStrKMP(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	lps := buildLPS(needle)
	i, j := 0, 0 // i duyệt haystack, j duyệt needle

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		}

		if j == len(needle) {
			return i - j // Trả về vị trí bắt đầu
		}

		if i < len(haystack) && haystack[i] != needle[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))

	fmt.Println(strStrKMP("sadbutsad", "sad"))
	fmt.Println(strStrKMP("leetcode", "leeto"))
}
