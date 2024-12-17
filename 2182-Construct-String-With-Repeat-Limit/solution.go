package main

import (
	"container/heap"
	"fmt"
)

// CharFrequency define a struct to store a character and its frequency
type CharFrequency struct {
	char rune
	freq int
}

// MaxHeap define a MaxHeap to maintain characters in descending order based on their values
type MaxHeap []CharFrequency

func (h MaxHeap) Len() int {
	return len(h)
}

// Less Max-heap based on character value
func (h MaxHeap) Less(i, j int) bool {
	return h[i].char > h[j].char
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// repeatLimitedString main function to generate the lexicographically largest repeat-limited string
func repeatLimitedString(s string, repeatLimit int) string {
	// Step 1: Count the frequency of each character
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	// Step 2: Push all characters with their frequencies into a MaxHeap
	h := &MaxHeap{}
	heap.Init(h)
	for ch, count := range freq {
		heap.Push(h, CharFrequency{char: ch, freq: count})
	}

	// Step 3: Build the result string
	var result []rune
	for h.Len() > 0 {
		// Extract the character with the highest value
		top := heap.Pop(h).(CharFrequency)
		char, count := top.char, top.freq

		// Use the character up to `repeatLimit` times
		times := min(count, repeatLimit)
		for i := 0; i < times; i++ {
			result = append(result, char)
		}

		// If there are remaining characters, we need to "break" the sequence
		if count > times {
			// If there are no more characters in the heap, break the loop
			if h.Len() == 0 {
				break
			}

			// Extract the next character from the heap
			next := heap.Pop(h).(CharFrequency)
			result = append(result, next.char)
			next.freq--

			// Push the next character back if it still has remaining frequency
			if next.freq > 0 {
				heap.Push(h, next)
			}

			// Push the current character (with reduced frequency) back into the heap
			heap.Push(h, CharFrequency{char: char, freq: count - times})
		}
	}

	return string(result)
}

func main() {
	fmt.Println(repeatLimitedString("cczazcc", 3)) // Output: "zzcccac"
	fmt.Println(repeatLimitedString("aababab", 2)) // Output: "bbabaa"
}
