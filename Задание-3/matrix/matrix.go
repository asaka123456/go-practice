package matrix

import (
	"math/rand"
	"time"
)

func GenerateUnique(rows, cols int) [][]int {
	rand.Seed(time.Now().UnixNano())
	total := rows * cols
	numbers := rand.Perm(total)

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[i*cols+j] + 1
		}
	}
	return matrix
}
