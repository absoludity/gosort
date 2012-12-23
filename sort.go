package sort

import (
	DefaultSort "sort"
)

func Ints(a []int) {
    Sort(DefaultSort.IntSlice(a))
}

func Sort(data DefaultSort.Interface) {
	var sorted bool
	for {
		sorted = true
		len_minus_1 := data.Len() - 1
		for i := 0; i < len_minus_1; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
				sorted = false
			}
		}
		if sorted {
			return
		}
	}
}
