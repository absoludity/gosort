package sort

import (
	"math/rand"
	DefaultSort "sort"
	"testing"
)

func makeRandomIntSlice(size int) RandomIntSlice {
	intSlice := make([]int, size)
	for i := 0; i < size; i++ {
		intSlice[i] = rand.Int()
	}
	buffer := make([]int, size)
	randomSlice := RandomIntSlice{&intSlice, &buffer}
	return randomSlice
}

type RandomIntSlice struct {
	origInts *[]int
	copyInts *[]int
}

func (ri *RandomIntSlice) GetCopy() *[]int {
	copy(*ri.copyInts, *ri.origInts)
	return ri.copyInts
}

var IntSlice10 RandomIntSlice = makeRandomIntSlice(10)

func BenchmarkDefaultSort10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultSort.Ints(*IntSlice10.GetCopy())
	}
}

func BenchmarkSort10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(*IntSlice10.GetCopy())
	}
}

var IntSlice100 RandomIntSlice = makeRandomIntSlice(100)

func BenchmarkDefaultSort100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultSort.Ints(*IntSlice100.GetCopy())
	}
}

func BenchmarkSort100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(*IntSlice100.GetCopy())
	}
}

var IntSlice1000 RandomIntSlice = makeRandomIntSlice(1000)

func BenchmarkDefaultSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultSort.Ints(*IntSlice1000.GetCopy())
	}
}

func BenchmarkSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ints(*IntSlice1000.GetCopy())
	}
}
