package main

import (
	"fmt"
	"math"
)

func main() {
	values := [][]float64{}
	// These are the first two rows.
	row1 := []float64{1, 2, 3}
	row2 := []float64{4, 8, 2}
	row3 := []float64{1, 5, 3}

	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	val := twoDArray(values, 3, 3)
	fmt.Println("Min cost is", val)
}
func twoDArray(val [][]float64, m, n int) float64 {
	temp := [3][3]float64{}
	sum := 0.0
	for i := 0; i < m; i++ {
		temp[0][i] = sum + val[0][i]
		sum = temp[0][i]
	}
	sum = 0.0
	for i := 0; i < n; i++ {
		temp[i][0] = sum + val[i][0]
		sum = temp[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			temp[i][j] = math.Min(temp[i][j-1], temp[i-1][j]) + val[i][j]
		}
	}
	return temp[m-1][n-1]
}
