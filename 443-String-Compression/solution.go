package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	idx, i := 0, 0

	for i < n {
		current := chars[i]
		count := 0

		for i < n && chars[i] == current {
			count++
			i++
		}

		chars[idx] = current
		idx++

		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[idx] = byte(digit)
				idx++
			}
		}
	}

	return idx
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a'}))
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}
