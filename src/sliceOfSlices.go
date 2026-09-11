package main

func createMatrix(rows, cols int) [][]int {
	// ?
	var matrix [][]int
	for i := 0; i < rows; i++{
		row := make([]int, cols)
		matrix = append(matrix, row)
		
		for j := 0; j < cols; j++{
			matrix[i][j] = i * j
		}
	}
	return matrix
}
