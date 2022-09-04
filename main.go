package main

import "fmt"

func AdjacencyList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		if _, OK := graph[edge[0]]; OK {
			graph[edge[0]] = append(graph[edge[0]], edge[1])
		} else {
			graph[edge[0]] = make([]int, 0)
			graph[edge[0]] = append(graph[edge[0]], edge[1])
		}
		if _, OK := graph[edge[1]]; OK {
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		} else {
			graph[edge[1]] = make([]int, 0)
			graph[edge[1]] = append(graph[edge[1]], edge[0])
		}
	}
	return graph
}
func DepthFirstSearch(graph map[int][]int, visited []bool, node int) {
	if visited[node] {
		return
	}
	fmt.Println(node)
	visited[node] = true
	for _, neighbor := range graph[node] {
		DepthFirstSearch(graph, visited, neighbor)
	}
}

func BreadthFirstSearch(graph map[int][]int, node int) {
	size := 11
	visited := make([]bool, size)
	queue := make([]int, 0)

	visited[node] = true
	queue = append(queue, node)
	for len(queue) > 0 {
		var value int
		value, queue = queue[0], queue[1:]
		fmt.Println(value)
		for _, neighbor := range graph[value] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
func main() {
	edges := [][]int{{0, 1}, {0, 5}, {0, 4}, {4, 2}, {2, 3}, {3, 9}}
	size := 11
	graph := AdjacencyList(edges)
	visited := make([]bool, size)
	start := 0
	DepthFirstSearch(graph, visited, start)
	BreadthFirstSearch(graph, 9)
}
