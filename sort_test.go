package sort

import (
	DefaultSort "sort"
	"testing"
)

func TestInts(t *testing.T) {
	a := []int{5, 7, 6, 4, 3}

	Ints(a)

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Errorf("Result not sorted:", a)
			break
		}
	}
}

func TestSort(t *testing.T) {
	a := []int{5, 7, 6, 4, 3}

	Sort(DefaultSort.IntSlice(a))

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			t.Errorf("Result not sorted:", a)
			break
		}
	}
}
