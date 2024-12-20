package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		n := len(queue)
		nodes := queue[:n]

		if level%2 == 1 {
			for i := 0; i < n/2; i++ {
				nodes[i].Val, nodes[n-i-1].Val = nodes[n-i-1].Val, nodes[i].Val
			}
		}

		for _, node := range nodes {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[n:]
		level++
	}

	return root
}

func createBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, val := range nums {
		nodes[i] = &TreeNode{Val: val}
	}

	// Liên kết các nút để tạo cây
	for i := 0; i < len(nums)/2; i++ {
		if 2*i+1 < len(nums) {
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < len(nums) {
			nodes[i].Right = nodes[2*i+2]
		}
	}

	return nodes[0]
}

func printTreeByLevel(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}

	fmt.Println("---------")
}

func main() {
	printTreeByLevel(reverseOddLevels(createBinaryTree([]int{2, 3, 5, 8, 13, 21, 34})))
	printTreeByLevel(reverseOddLevels(createBinaryTree([]int{7, 13, 11})))
	printTreeByLevel(reverseOddLevels(createBinaryTree([]int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2})))
}
