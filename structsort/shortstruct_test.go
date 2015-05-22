package structsort

import (
	"math/rand"
	"sort"
	"testing"
)

type TS struct {
	value  uint64
	v1, v2 uint32
	flag   bool
}

func (x *TS) setRandom() {
	x.value = uint64(rand.Int63())

	// we do not really care what's inside but let's do this:
	x.flag = x.value%2 == 0
	x.v1 = uint32(x.value % 3)
	x.v2 = uint32(x.value % 5)
}

type shortptr []*TS

func (x shortptr) Len() int           { return len(x) }
func (x shortptr) Less(i, j int) bool { return x[i].value < x[j].value }
func (x shortptr) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type shortplain []TS

func (x shortplain) Len() int           { return len(x) }
func (x shortplain) Less(i, j int) bool { return x[i].value < x[j].value }
func (x shortplain) Swap(i, j int) {
	x[i].value, x[j].value = x[j].value, x[i].value
	x[i].flag, x[j].flag = x[j].flag, x[i].flag
	x[i].v1, x[j].v1 = x[j].v1, x[i].v2
	x[i].v2, x[j].v2 = x[j].v2, x[i].v1
}

func BenchmarkShortStructSortPtr(b *testing.B) {
	l := shortptr(make([]*TS, b.N))
	for i := 0; i < b.N; i++ {
		l[i] = new(TS)
		l[i].setRandom()
	}
	b.ResetTimer()
	sort.Sort(l)
}

func BenchmarkShortStructSortPlain(b *testing.B) {
	l := shortplain(make([]TS, b.N))
	for i := 0; i < b.N; i++ {
		l[i].setRandom()
	}
	b.ResetTimer()
	sort.Sort(l)
}

func BenchmarkShortStructFillPtr(b *testing.B) {
	l := shortptr(make([]*TS, b.N))
	for i := 0; i < b.N; i++ {
		l[i] = new(TS)
		l[i].setRandom()
	}
}

func BenchmarkShortStructFillPlain(b *testing.B) {
	l := shortplain(make([]TS, b.N))
	for i := 0; i < b.N; i++ {
		l[i].setRandom()
	}
}
