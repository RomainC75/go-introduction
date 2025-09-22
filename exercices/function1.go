package exercices

import "fmt"

func RunFunction1() {
	slc := []int{3, 5, 4, 9, 6, 34, 65, 213}

	res := filter(slc, func(value int) bool {
		return value < 10
	})

	fmt.Println(res)
}

func filter(slc []int, fn func(int) bool) []int {
	res := make([]int, 0, len(slc))
	for _, value := range slc {
		if fn(value) {
			res = append(res, value)
		}
	}
	return res
}
