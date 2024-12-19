package main

import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}

		count := 0
		for i >= 0 && s[i] != ' ' {
			count++
			i--
		}

		return count
	}

	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Solution goes here!"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("     "))
}
