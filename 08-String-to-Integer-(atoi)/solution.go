package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	charMap := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}
	sign, res, i := 1, 0, 0
	n := len(s)
	for i < n && s[i] == ' ' {
		i++
	}

	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < n {
		if num, exist := charMap[s[i]]; exist {
			res = 10*res + num
		} else {
			break
		}

		if res*sign < math.MinInt32 {
			return math.MinInt32
		}
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}

		i++
	}

	return res * sign
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -042"))
	fmt.Println(myAtoi("1337c0d3"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("words and 987"))
}
