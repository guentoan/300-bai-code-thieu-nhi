package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count, zeroCount := 0, 1
	for _, f := range flowerbed {
		if f == 0 {
			zeroCount++
		} else {
			count += (zeroCount - 1) / 2
			zeroCount = 0
		}
	}

	count += zeroCount / 2

	return count >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
