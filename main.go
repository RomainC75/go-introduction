package main

import "fmt"

func main() {
	matrix := [4][4]int{}
	i, j := 0, len(matrix[0])-1
	fmt.Println(i, j)
	for i < len(matrix) {
		matrix[i][j] = 1
		i++
		j--
	}
	displayMatrix(matrix)
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
