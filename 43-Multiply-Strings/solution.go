package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func padZeros(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}

func multiply(x, y string) string {
	// If either x or y is a single digit, directly multiply
	if len(x) == 1 || len(y) == 1 {
		n1, _ := strconv.Atoi(x)
		n2, _ := strconv.Atoi(y)
		return strconv.Itoa(n1 * n2)
	}

	// Find the length of the largest number
	m := int(math.Max(float64(len(x)), float64(len(y))))
	m2 := m / 2 // Split point for Karatsuba
	x = padZeros(x, m)
	y = padZeros(y, m)

	// Split both numbers into high and low parts
	x1, x0 := x[:len(x)-m2], x[len(x)-m2:]
	y1, y0 := y[:len(y)-m2], y[len(y)-m2:]

	// Recursively calculate the three parts of Karatsuba
	z2 := multiply(x1, y1)                   // High part multiplication: a * c
	z0 := multiply(x0, y0)                   // Low part multiplication: b * d
	z1 := multiply(add(x1, x0), add(y1, y0)) // (a + b) * (c + d)

	// Combine the results according to Karatsuba formula
	result := add(add(shift(z2, 2*m2), shift(subtract(z1, add(z2, z0)), m2)), z0)

	return trimLeadingZeros(result)
}

// add function to add two large numbers represented as strings
func add(x, y string) string {
	// Ensure x is the larger number
	if len(x) < len(y) {
		x, y = y, x
	}

	// Perform addition digit by digit, handling carry
	carry := 0
	result := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		a := int(x[len(x)-i-1] - '0')
		b := 0
		if i < len(y) {
			b = int(y[len(y)-i-1] - '0')
		}
		sum := a + b + carry
		result[len(x)-i-1] = byte(sum%10 + '0')
		carry = sum / 10
	}

	// If carry is left over, prepend it to the result
	if carry > 0 {
		return string(append([]byte{byte(carry + '0')}, result...))
	}
	return string(result)
}

// subtract function to subtract two large numbers represented as strings
func subtract(x, y string) string {
	// If both numbers are equal, return "0"
	if x == y {
		return "0"
	}

	// Perform subtraction digit by digit, handling borrow
	carry := 0
	result := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		a := int(x[len(x)-i-1] - '0')
		b := 0
		if i < len(y) {
			b = int(y[len(y)-i-1] - '0')
		}
		sub := a - b - carry
		if sub < 0 {
			sub += 10
			carry = 1
		} else {
			carry = 0
		}
		result[len(x)-i-1] = byte(sub + '0')
	}
	return strings.TrimLeft(string(result), "0")
}

func trimLeadingZeros(num string) string {
	i := 0
	for i < len(num) && num[i] == '0' {
		i++
	}
	if i == len(num) {
		return "0"
	}
	return num[i:]
}

// shift function to shift the result by n digits (equivalent to multiplying by 10^n)
func shift(x string, n int) string {
	// Append n zeros to the number
	return x + strings.Repeat("0", n)
}

func main() {
	fmt.Println(multiply("2", "3"))     // 6
	fmt.Println(multiply("123", "456")) // 56088
	fmt.Println(multiply("3", "33"))    // 99
	fmt.Println(multiply("9", "99"))    // 891
	fmt.Println(multiply("6994", "36")) // 251784
}
