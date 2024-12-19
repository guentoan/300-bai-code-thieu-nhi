package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(12012121222999999))
}
