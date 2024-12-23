package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	count := 0
	for len(queue) > 0 {
		n := len(queue)
		nodes := queue[:n]
		values := make([]int, n)
		for i, node := range nodes {
			values[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count += minSwap(values)
		queue = queue[n:]
	}

	return count
}

func minSwap(values []int) int {
	n := len(values)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{values[i], i}
	}

	// Sort pairs based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		if visited[i] || pairs[i][1] == i {
			continue
		}

		// Cycle detection
		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pairs[j][1]
			cycleSize++
		}

		// If cycle size > 1, add (cycleSize - 1) swaps
		if cycleSize > 1 {
			swaps += cycleSize - 1
		}
	}

	return swaps
}

func createBinaryTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root} // Queue for BFS
	index := 1

	for len(queue) > 0 && index < len(values) {
		current := queue[0]
		queue = queue[1:] // Dequeue

		// Add left child
		if index < len(values) && values[index] != nil {
			current.Left = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Left)
		}
		index++

		// Add right child
		if index < len(values) && values[index] != nil {
			current.Right = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

func main() {
	fmt.Println(minimumOperations(createBinaryTree([]interface{}{1, 4, 3, 7, 6, 8, 5, nil, nil, nil, nil, 9, nil, 10})))
	fmt.Println(minimumOperations(createBinaryTree([]interface{}{1, 3, 2, 7, 6, 5, 4})))
	fmt.Println(minimumOperations(createBinaryTree([]interface{}{1, 2, 3, 4, 5, 6})))
}
