package main

import (
	"container/heap"
	"fmt"
	"math"
)

// MaxHeap implementation
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, gift := range gifts {
		heap.Push(h, gift)
	}

	for i := 0; i < k; i++ {
		maxVal := heap.Pop(h).(int)
		newVal := int(math.Floor(math.Sqrt(float64(maxVal))))
		heap.Push(h, newVal)
	}

	var ans int64
	for h.Len() > 0 {
		ans += int64(heap.Pop(h).(int))
	}

	return ans
}

func main() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4)) // expected 29
	fmt.Println(pickGifts([]int{1, 1, 1, 1}, 4))        // expected 4
}
