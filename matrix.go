package matrix

import "fmt"

type Matrix [][]int

func NewMatrix(n, m int) Matrix {
	matrix := make(Matrix, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	return matrix
}

func (m Matrix) Scale(c int) {
	for i, row := range m {
		for j, entry := range row {
			m[i][j] = entry * c
			fmt.Printf("%d %d: %d\n", i, j, entry)
		}
	}
}
