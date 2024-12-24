package main

import "fmt"

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	n, m := len(edges1)+1, len(edges2)+1
	graph1, graph2 := make([][]int, n), make([][]int, m)

	for _, edge := range edges1 {
		graph1[edge[0]] = append(graph1[edge[0]], edge[1])
		graph1[edge[1]] = append(graph1[edge[1]], edge[0])
	}

	for _, edge := range edges2 {
		graph2[edge[0]] = append(graph2[edge[0]], edge[1])
		graph2[edge[1]] = append(graph2[edge[1]], edge[0])
	}

	d1 := getDiameter(graph1, n)
	d2 := getDiameter(graph2, m)
	cand := (d1+1)/2 + (d2+1)/2 + 1

	return max(d1, d2, cand)
}

func getDiameter(graph [][]int, n int) int {
	// Function to perform DFS and calculate the farthest node and its distance
	var dfs func(node int, visited []bool) (int, int)
	dfs = func(node int, visited []bool) (int, int) {
		visited[node] = true
		maxDist := 0
		farthestNode := node

		// Traverse all neighbors
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dist, farthest := dfs(neighbor, visited)
				if dist+1 > maxDist {
					maxDist = dist + 1
					farthestNode = farthest
				}
			}
		}
		return maxDist, farthestNode
	}

	// Step 1: Run DFS from any node (e.g., node 0) to find the farthest node
	visited := make([]bool, n)
	_, farthestNode := dfs(0, visited)

	// Step 2: Run DFS again from the farthest node to find the maximum distance
	visited = make([]bool, n)
	diameter, _ := dfs(farthestNode, visited)

	return diameter
}

func main() {
	fmt.Println(minimumDiameterAfterMerge([][]int{{0, 1}, {0, 2}, {0, 3}}, [][]int{{0, 1}}))
	fmt.Println(minimumDiameterAfterMerge([][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}))
}
