package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentVal := romanMap[s[i]]
		if i < n-1 && romanMap[s[i+1]] > currentVal {
			ans -= currentVal
		} else {
			ans += currentVal
		}
	}

	return ans
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
