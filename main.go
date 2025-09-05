package main

import "fmt"

func main() {

	slc1 := []int{1, 2, 3, 4, 5}
	slc2 := make([]int, 4, 5)
	fmt.Printf("slc1 : %p / len : %v / cap : %v \n", slc1, len(slc1), cap(slc1))
	fmt.Printf("PRE slc2 : %p / len : %v / cap : %v \n", slc2, len(slc2), cap(slc2))

	nbCopied := copy(slc2, slc1)
	fmt.Println("copied : ", nbCopied)
	fmt.Printf("POST slc2 : %p / len : %v / cap : %v \n", slc2, len(slc2), cap(slc2))
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
