package main

import "fmt"

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true,
		'i': true,
		'u': true,
		'e': true,
		'o': true,
		'A': true,
		'I': true,
		'U': true,
		'E': true,
		'O': true,
	}

	arr := []rune(s)
	left, right := 0, len(arr)-1
	for left < right {
		for left < right && !vowels[arr[left]] {
			left++
		}

		for left < right && !vowels[arr[right]] {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return string(arr)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
	fmt.Println(reverseVowels("leetcode"))
}
