package main

import "fmt"

func main() {
	slc := make([]int, 0, 0)
	// fmt.Printf("SLC : value : %d // len : %v // cap : %v\n", i, len(slc), cap(slc))

	for i := 0; i < 1_000_000; i++ {
		slc = append(slc, 0)
		if i < 100 {
			fmt.Printf("SLC : value : %d // len : %v // cap : %v // address : %p \n", i, len(slc), cap(slc), slc)
		}
		if i < 10_000 && i%100 == 0 {
			fmt.Printf("SLC : value : %d // len : %v // cap : %v // address : %p \n", i, len(slc), cap(slc), slc)
		}
		if i%10_000 == 0 {
			fmt.Printf("SLC : value : %d // len : %v // cap : %v // address : %p \n", i, len(slc), cap(slc), slc)
		}
	}
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
