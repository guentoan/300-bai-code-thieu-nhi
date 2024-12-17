package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum, current, startAt := 0, 0, 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		current += gas[i] - cost[i]
		if current < 0 {
			startAt = i + 1
			current = 0
		}
	}

	if sum >= 0 {
		return startAt
	}

	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
