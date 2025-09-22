package mergesort_test

import (
	mergesort "sandbox/merge_sort"
	"testing"
)

// run :
// cd merge_sort
// go test -bench=.

var slc1 = []int{1, 4, 6, 23, 54, 100, 110, 150, 210, 315, 654, 734, 769}
var slc2 = []int{3, 6, 32, 115, 342, 350, 351, 435, 2432}

func BenchmarkStandard(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergesort.SortStandard(slc1, slc2)
	}
}

func BenchmarkAlgo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergesort.SortAlgo(slc1, slc2)
	}
}
