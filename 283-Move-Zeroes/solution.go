package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}

		j := i + 1
		for j < len(nums) {
			if nums[j] != 0 {
				break
			}
			j++
		}

		if j == len(nums) {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
