package slices

import (
	"testing"
)

// let's try to see if it's fast to re-assign slice var

func BenchmarkSliceReassign(b *testing.B) {
	awkward := make([]byte, b.N+4)
	for i := 0; i < b.N; i++ {
		awkward := awkward[1:]
		test1, test2, test3, test4 := awkward[0], awkward[1], awkward[2], awkward[3] // assume decoding uint32 from those
		_, _, _, _ = test1, test2, test3, test4
	}
}

func BenchmarkSliceMove(b *testing.B) {
	awkward := make([]byte, b.N+4)
	pos := 0
	for i := 0; i < b.N; i++ {
		test1, test2, test3, test4 := awkward[pos+1], awkward[pos+2], awkward[pos+3], awkward[pos+4] // about same work
		pos++
		_, _, _, _ = test1, test2, test3, test4
	}
}
