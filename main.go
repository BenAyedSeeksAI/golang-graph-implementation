package main

import "fmt"

func LoopThrough(graph map[int][]int, node int, visited []bool) bool {
	if visited[node] {
		return false
	}
	visited[node] = true
	for _, neighbor := range graph[node] {
		LoopThrough(graph, neighbor, visited)
	}
	return true
}

func ConnectedComponents(graph map[int][]int, size int) int {
	visited := make([]bool, size)
	var cnt int = 0
	for node := range graph {
		if LoopThrough(graph, node, visited) == true {
			cnt++
		}
	}
	return cnt
}

func AjacencyMatrix(edges [][]int, n int) [][]int {
	matrix := make([][]int, n)
	for idx := range matrix {
		matrix[idx] = make([]int, n)
	}
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = 1
		matrix[edge[1]][edge[0]] = 1
	}
	return matrix
}

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
func graph() {
	edges := [][]int{{0, 1}, {0, 5}, {0, 4}, {2, 6}, {6, 3}, {3, 9}}
	size := 10
	graphLst := AdjacencyList(edges)
	graphMtx := AjacencyMatrix(edges, size)
	fmt.Println("Adjacency List :")
	fmt.Println(graphLst)
	fmt.Println("Adjacency Matrix :")
	fmt.Println(graphMtx)
	fmt.Println(ConnectedComponents(graphLst, size))
	visited := make([]bool, size)
	start := 0
	DepthFirstSearch(graphLst, visited, start)
	BreadthFirstSearch(graphLst, 9)
}
func main() {
	// graph()
	mainTree()
}
