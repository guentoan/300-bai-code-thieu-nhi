package main

import (
	"fmt"
	"strings"
)

func slidingPuzzle(board [][]int) int {
	// The target solved state
	target := "123450"

	// Convert the 2D board to a string representation
	start := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start += fmt.Sprint(board[i][j])
		}
	}

	// If the initial state is already solved, return 0
	if start == target {
		return 0
	}

	// Define the possible moves for each position of '0'
	neighbors := [][]int{
		{1, 3},    // position 0 can move with positions 1 and 3
		{0, 2, 4}, // position 1 can move with positions 0, 2, and 4
		{1, 5},    // position 2 can move with positions 1 and 5
		{0, 4},    // position 3 can move with positions 0 and 4
		{1, 3, 5}, // position 4 can move with positions 1, 3, and 5
		{2, 4},    // position 5 can move with positions 2 and 4
	}

	// BFS
	queue := []string{start}                // Initialize the queue with the starting state
	visited := map[string]bool{start: true} // Set to keep track of visited states
	steps := 0                              // Variable to count the number of moves

	// While there are states to process in the queue
	for len(queue) > 0 {
		var nextQueue []string // Temporary queue to hold next states

		// Process each state in the current queue
		for _, curr := range queue {
			// Find the index of '0' in the current state string
			zeroIndex := strings.Index(curr, "0")

			// Try all possible moves for the current '0'
			for _, neighbor := range neighbors[zeroIndex] {
				// Swap '0' with the adjacent number to create a new state
				newBoard := []rune(curr)
				newBoard[zeroIndex], newBoard[neighbor] = newBoard[neighbor], newBoard[zeroIndex]
				newState := string(newBoard)

				// If the new state matches the target, return the number of steps
				if newState == target {
					return steps + 1
				}

				// If the new state has not been visited, add it to the next queue
				if !visited[newState] {
					visited[newState] = true
					nextQueue = append(nextQueue, newState)
				}
			}
		}

		// Move to the next level (bigger step count)
		queue = nextQueue
		steps++
	}

	// If no solution is found, return -1
	return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{
		{1, 2, 3},
		{4, 0, 5},
	})) // Expected 1

	fmt.Println(slidingPuzzle([][]int{
		{1, 2, 3},
		{5, 4, 0},
	})) // Expected -1

	fmt.Println(slidingPuzzle([][]int{
		{4, 1, 2},
		{5, 0, 3},
	})) // Expected 5
}
