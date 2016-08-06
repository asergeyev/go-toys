package pvsi

import (
	"testing"
)

type data struct {
	b [1024]byte
	s string // ""
}

const LEN = 10000

var store [LEN]data

func allocatePtrs() []*data {
	p := make([]*data, len(store))
	for i := range p {
		p[i] = &(store[i])
	}
	return p
}

func allocateIfs() []interface{} {
	p := make([]interface{}, len(store))
	for i := range p {
		p[i] = &(store[i])
	}
	return p
}

func processPtrs(p []*data) byte {
	var sum byte
	for _, v := range p {
		sum += v.b[512]
	}
	return sum
}

func processIfs(p []interface{}) byte {
	var sum byte
	for _, v := range p {
		sum += v.(*data).b[512]
	}
	return sum
}

func TestPtrs(t *testing.T) {
	x := allocatePtrs()
	s := processPtrs(x)
	if s != 0 {
		t.Error("Invalid result")
	}
}

func TestIfs(t *testing.T) {
	x := allocateIfs()
	s := processIfs(x)
	if s != 0 {
		t.Error("Invalid result")
	}
}

func BenchmarkAllocatePtrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := allocatePtrs()
		if x == nil {
			b.Fatal("Error allocating ptrs")
		}
	}
}

func BenchmarkAllocateIfs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := allocateIfs()
		if x == nil {
			b.Fatal("Error allocating ifs")
		}
	}
}

func BenchmarkProcessPtrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x := allocatePtrs()
		b.StartTimer()
		s := processPtrs(x)
		if s != 0 {
			b.Fatal("Error processing ptrs")
		}
	}
}

func BenchmarkProcessIfs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x := allocateIfs()
		b.StartTimer()
		s := processIfs(x)
		if s != 0 {
			b.Fatal("Error processing ifs")
		}
	}
}
