package exercices

import "fmt"

func RunFunction1() {
	slc := []int{3, 5, 4, 9, 6, 34, 65, 213}

	res := filter(slc, func(value int) bool {
		return value < 10
	})

	fmt.Println(res)
}

type Int interface {
	int8 | uint | int
}

func filter[T Int](slc []T, fn func(T) bool) []T {
	res := make([]T, 0, len(slc))
	for _, value := range slc {
		if fn(value) {
			res = append(res, value)
		}
	}
	return res
}
