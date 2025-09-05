package main

import "fmt"

func main() {

	slc1 := []int{1, 2, 3, 4}
	fmt.Printf("%p : len : %d // cap : %d \n", slc1, len(slc1), cap(slc1))

	slc1 = append(slc1[:2], slc1[3:]...)
	fmt.Printf("%p : len : %d // cap : %d \n", slc1, len(slc1), cap(slc1))
	fmt.Println(slc1)
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
