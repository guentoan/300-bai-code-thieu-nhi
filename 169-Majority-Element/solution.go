package main

import "fmt"

func majorityElement(nums []int) int {
	var ans, count int
	for _, n := range nums {
		if count == 0 {
			ans = n
			count++
			continue
		}

		if ans == n {
			count++
		} else {
			count--
		}
	}

	return ans
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
