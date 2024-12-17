package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	i := n - 1
	result := make([]string, 0)
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}

		j := i
		for i >= 0 && s[i] != ' ' {
			i--
		}

		if j >= 0 {
			result = append(result, s[i+1:j+1])
		}
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
