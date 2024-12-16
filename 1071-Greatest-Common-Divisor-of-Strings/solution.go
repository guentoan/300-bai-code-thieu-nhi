package main

import "fmt"

func gdc(a int, b int) int {
	if b == 0 {
		return a
	}

	return gdc(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	x := gdc(len(str1), len(str2))

	return str1[:x]
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "AB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}
