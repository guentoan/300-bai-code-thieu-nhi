package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	result := "1"
	for i := 2; i <= n; i++ {
		result = runLengthEncoding(result)
	}

	return result
}

func runLengthEncoding(s string) string {
	builder := strings.Builder{}
	count := 1
	for i := 1; i < len(s); i++ {
		// If the current character is the same as the previous one, increment the count
		if s[i] == s[i-1] {
			count++
		} else {
			// When the character changes, append the count and the character to the nextResult
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(s[i-1])
			// Reset count to 1 for the new character
			count = 1
		}
	}
	builder.WriteString(strconv.Itoa(count))
	builder.WriteByte(s[len(s)-1])
	return builder.String()
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(4))
}
