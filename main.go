package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slc := arr[1:]
	fmt.Println(arr)
	fmt.Printf("SLC : value : %v // len : %v // cap : %v\n", slc, len(slc), cap(slc))

	slc[1] = 34
	fmt.Println(arr)
	fmt.Printf("SLC : value : %v // len : %v // cap : %v\n", slc, len(slc), cap(slc))

	slc2 := slc[2:4]
	slc2[0] = 48
	fmt.Printf("SLC2 : value : %v // len : %v // cap : %v\n", slc2, len(slc2), cap(slc2))
	fmt.Println("final", arr)
}

func displayMatrix(matrix [4][4]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
