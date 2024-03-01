package main

import (
	"fmt"
)

func generarTrianguloPascal(n int) [][]int {
	pascal := make([][]int, n)
	for i := range pascal {
		pascal[i] = make([]int, i+1)
		pascal[i][0], pascal[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
	}
	return pascal
}

func main() {
	nivel := 8
	pascal := generarTrianguloPascal(nivel)

	fmt.Println("Triángulo de Pascal (primeras 8 filas):")
	for i := 0; i < nivel; i++ {
		fmt.Println(pascal[i])
	}
}
