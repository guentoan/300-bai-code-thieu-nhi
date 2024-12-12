package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	result := make([]byte, len1+len2)
	i, j := 0, 0

	for j < len1 && j < len2 {
		result[i] = word1[j]
		result[i+1] = word2[j]
		j++
		i += 2
	}

	for j < len1 {
		result[i] = word1[j]
		j++
		i++
	}
	for j < len2 {
		result[i] = word2[j]
		j++
		i++
	}

	return string(result)
}

func mergeAlternatelyStringBuilder(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	var result strings.Builder
	maxLen := max(len1, len2)
	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result.WriteByte(word1[i])
		}
		if i < len2 {
			result.WriteByte(word2[i])
		}
	}

	return result.String()
}

func mergeAlternatelyString(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	var result string
	maxLen := max(len1, len2)
	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result += string(word1[i])
		}
		if i < len2 {
			result += string(word2[i])
		}
	}

	return result
}

func main() {
	fmt.Println(mergeAlternately("banana", "apple"))
	fmt.Println(mergeAlternately("abcd", "pq"))

	fmt.Println(mergeAlternatelyStringBuilder("banana", "apple"))
	fmt.Println(mergeAlternatelyStringBuilder("abcd", "pq"))

	fmt.Println(mergeAlternatelyString("banana", "apple"))
	fmt.Println(mergeAlternatelyString("abcd", "pq"))
}
