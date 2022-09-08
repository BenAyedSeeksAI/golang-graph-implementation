package main

func initMatrix() [][]int {
	matrix := make([][]int, 11)
	for idx := range matrix {
		matrix[idx] = make([]int, 11)
	}
	return matrix
}
