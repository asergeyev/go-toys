package structsort

import (
	"math/rand"
	"sort"
	"testing"
)

const ALEN = 20

type TL struct {
	value uint64
	array [ALEN]uint32 // to make it more complex to copy
	flag  bool
}

func (x *TL) setRandom() {
	x.value = uint64(rand.Int63())

	// we do not really care what's inside but let's do this:
	x.flag = x.value%2 == 0
	for i := range x.array {
		x.array[i] = uint32(i)
	}
}

type longptr []*TL

func (x longptr) Len() int           { return len(x) }
func (x longptr) Less(i, j int) bool { return x[i].value < x[j].value }
func (x longptr) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type longplain []TL

func (x longplain) Len() int           { return len(x) }
func (x longplain) Less(i, j int) bool { return x[i].value < x[j].value }
func (x longplain) Swap(i, j int) {
	x[i].value, x[j].value = x[j].value, x[i].value
	x[i].flag, x[j].flag = x[j].flag, x[i].flag

	var tmp [ALEN]uint32
	copy(tmp[:], x[i].array[:])
	copy(x[i].array[:], x[j].array[:])
	copy(x[j].array[:], tmp[:])
}

func BenchmarkLongStructSortPtr(b *testing.B) {
	l := longptr(make([]*TL, b.N))
	for i := 0; i < b.N; i++ {
		l[i] = new(TL)
		l[i].setRandom()
	}
	b.ResetTimer()
	sort.Sort(l)
}

func BenchmarkLongStructSortPlain(b *testing.B) {
	l := longplain(make([]TL, b.N))
	for i := 0; i < b.N; i++ {
		l[i].setRandom()
	}
	b.ResetTimer()
	sort.Sort(l)
}

func BenchmarkLongStructFillPtr(b *testing.B) {
	l := longptr(make([]*TL, b.N))
	for i := 0; i < b.N; i++ {
		l[i] = new(TL)
		l[i].setRandom()
	}
}

func BenchmarkLongStructFillPlain(b *testing.B) {
	l := longplain(make([]TL, b.N))
	for i := 0; i < b.N; i++ {
		l[i].setRandom()
	}
}
