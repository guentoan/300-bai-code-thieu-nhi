package main

import (
	"fmt"
	"sort"
)

const (
	startTime = 1
	endTime   = 0
)

func maxTwoEvents(events [][]int) int {
	times := make([][]int, 0)
	for _, event := range events {
		times = append(times, []int{event[0], startTime, event[2]})
		times = append(times, []int{event[1] + 1, endTime, event[2]})
	}

	sort.Slice(times, func(i, j int) bool {
		if times[i][0] == times[j][0] {
			return times[i][1] < times[j][1]
		}
		return times[i][0] < times[j][0]
	})

	maxValue, res := 0, 0
	for _, time := range times {
		if time[1] == startTime {
			res = max(res, time[2]+maxValue)
		} else {
			maxValue = max(maxValue, time[2])
		}
	}

	return res
}

func main() {
	events := [][]int{
		{1, 3, 2},
		{4, 5, 2},
		{2, 4, 3},
	}

	result := maxTwoEvents(events)
	fmt.Println(result)
}
