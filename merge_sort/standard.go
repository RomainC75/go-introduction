package mergesort

import "slices"

func SortStandard(slc1 []int, slc2 []int) []int {
	res := make([]int, 0, len(slc1)+len(slc2))
	res = append(res, slc1...)
	res = append(res, slc2...)

	slices.SortFunc(res, func(a, b int) int {
		return a - b
	})
	return res
}
