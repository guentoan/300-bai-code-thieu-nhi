package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	for _, n := range nums {
		if n != val {
			nums[j] = n
			j++
		}
	}

	return j
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
