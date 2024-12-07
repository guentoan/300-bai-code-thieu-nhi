package main

import (
	"fmt"
	"math"
)

func solution(nums []int, maxOperations int) int {
	canDivide := func(p int) bool {
		operations := 0
		for _, n := range nums {
			if n > p {
				operations += int(math.Ceil(float64(n)/float64(p))) - 1
			}
			if operations > maxOperations {
				return false
			}
		}
		return true
	}

	low, high := 1, 0
	for _, n := range nums {
		if n > high {
			high = n
		}
	}

	for low < high {
		mid := (low + high) / 2
		if canDivide(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func main() {
	// expected 3
	pen := solution([]int{9}, 2)
	fmt.Println(pen)

	// expected 2
	pen = solution([]int{2, 4, 8, 2}, 4)
	fmt.Println(pen)

	// expected 129
	pen = solution([]int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516,
		595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}, 80)
	fmt.Println(pen)
}
