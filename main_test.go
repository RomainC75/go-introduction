package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	slc1 := generateSlc(1000, 5, 6)
	slc2 := generateSlc(1000, 3, 32)
	fmt.Println(slc1)
	fmt.Println(slc2)
}

func generateSlc(max int, jumps ...int) []int {
	slc := make([]int, 0)
	for i := 0; i < max; i++ {
		if isDivisible(i, jumps) {
			slc = append(slc, i)
		}
	}
	return slc
}

func isDivisible(val int, jumps []int) bool {
	for _, jump := range jumps {
		if val%jump == 0 {
			return true
		}
	}
	return false
}
