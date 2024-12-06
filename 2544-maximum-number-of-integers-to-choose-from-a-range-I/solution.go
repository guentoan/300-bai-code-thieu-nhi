package main

import "fmt"

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]bool)
	for _, i := range banned {
		bannedMap[i] = true
	}

	count, sum := 0, 0
	for i := 1; i <= n; i++ {
		if !bannedMap[i] {
			if sum+i <= maxSum {
				sum += i
				count++
			} else {
				break
			}
		}
	}

	return count
}

func main() {
	banned := []int{1, 6, 5}
	n := 5
	maxSum := 6

	res := maxCount(banned, n, maxSum)
	fmt.Println(res)
}
