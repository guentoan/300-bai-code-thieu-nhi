package main

import (
	"container/heap"
	"fmt"
)

type Class struct {
	passed, total int
	improvement   float64
}

type MaxHeap []*Class

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].improvement > h[j].improvement }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*Class)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func calculateImprovement(passed, total int) float64 {
	current := float64(passed) / float64(total)
	newPassRatio := float64(passed+1) / float64(total+1)
	return newPassRatio - current
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, c := range classes {
		passed, total := c[0], c[1]
		heap.Push(h, &Class{
			passed:      passed,
			total:       total,
			improvement: calculateImprovement(passed, total),
		})
	}

	for extraStudents > 0 {
		top := heap.Pop(h).(*Class)
		top.passed++
		top.total++
		top.improvement = calculateImprovement(top.passed, top.total)
		heap.Push(h, top)
		extraStudents--
	}

	totalRatio := 0.0
	for h.Len() > 0 {
		c := heap.Pop(h).(*Class)
		totalRatio += float64(c.passed) / float64(c.total)
	}

	return totalRatio / float64(len(classes))
}

func main() {
	fmt.Println(maxAverageRatio([][]int{
		{1, 2},
		{3, 5},
		{2, 2},
	}, 2))
	fmt.Println(maxAverageRatio([][]int{
		{2, 4},
		{3, 9},
		{4, 5},
		{2, 10},
	}, 4))
}
