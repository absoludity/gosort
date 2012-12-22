package sort

import (
	DefaultSort "sort"
	"testing"
)

var ShortSlice []int = []int{5, 2, 6, 3, 1, 4}
var LongSlice []int = []int{
1, 1, 10, 8, 3, 6, 7, 4, 3, 8, 2, 9, 10, 7, 7, 4, 9, 8, 1, 9, 10, 9, 10, 2, 5, 7, 3, 8, 6, 4, 8, 3, 6, 10, 8, 2, 1, 10, 10, 2, 4, 8, 3, 7, 8, 6, 3, 10, 2, 7, 3, 5, 6, 2, 6, 10, 8, 1, 7, 10, 5, 6, 8, 5, 7, 6, 8, 9, 5, 7, 2, 9, 10, 6, 7, 8, 10, 3, 1, 7, 5, 9, 1, 8, 5, 1, 9, 10, 3, 8, 8, 10, 10, 8, 4, 8, 4, 7, 2, 3,
}

func BenchmarkDefaultIntsSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
        DefaultSort.Ints(ShortSlice)
	}
}

func BenchmarkDefaultIntsLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
        DefaultSort.Ints(LongSlice)
	}
}

func BenchmarkIntsShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Ints(ShortSlice)
	}
}

func BenchmarkIntsLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Ints(LongSlice)
	}
}
