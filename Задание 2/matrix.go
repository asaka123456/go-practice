package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateUniqueMatrix(rows, cols int) [][]int {
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

func main() {
	rand.Seed(time.Now().UnixNano())

	rows, cols := 5, 5
	fmt.Printf("Уникальная матрица %dx%d:\n\n", rows, cols)

	matrix := generateUniqueMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}
}
