package main

import "fmt"

func main() {

	slc1 := []int{1, 2, 3}
	fmt.Printf("%p\n", slc1)
	slc2 := []int{4, 5, 6}

	slc1 = append(slc1[:2], slc2[1:]...)
	fmt.Printf("%p\n", slc1)
	fmt.Println(slc1)
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
