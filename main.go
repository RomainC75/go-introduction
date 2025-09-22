package main

import "fmt"

func main() {
	slc1 := []int{1, 4, 6, 23, 54, 100, 110}
	slc2 := []int{3, 6, 32, 115, 342}
	res := make([]int, 0, len(slc1)+len(slc2))

	i1 := 0
	i2 := 0

	for i1 < len(slc1)-1 && i2 < len(slc2)-1 {
		if slc1[i1] < slc2[i2] {
			res = append(res, slc1[i1])
			i1++
		} else {
			res = append(res, slc2[i2])
			i2++
		}
	}

	if i1 < len(slc1)-1 {
		res = append(res, slc1[i1:]...)
	}
	if i2 < len(slc2)-1 {
		res = append(res, slc2[i2:]...)
	}
	fmt.Println(res)
}
