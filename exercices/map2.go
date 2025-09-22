package exercices

import (
	"fmt"
	"slices"
)

func Occurences() {
	slc := []int{0, 1, 4, 3, 6, 4, 3, 8, 5, 4, 3, 7, 5, 3, 6}

	occurences := map[int]int{}

	for _, value := range slc {
		_, ok := occurences[value]
		if ok {
			occurences[value]++
		} else {
			occurences[value] = 1
		}
	}

	slc2 := [][]int{}
	for k, v := range occurences {
		slc2 = append(slc2, []int{k, v})
	}
	slices.SortFunc(slc2, func(a []int, b []int) int {
		return a[1] - b[1]
	})

	fmt.Println(slc2)

	res := []int{}
	for _, v := range slc2 {
		res = append(res, v[0])
	}
	fmt.Println(res)
}
