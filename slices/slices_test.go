package slices

import (
	"testing"
)

// let's try to see if it's fast to re-assign slice var (spoiler alert: it is)

var testbuf = make([]byte, 8192)

func BenchmarkSliceReassign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		awkward := testbuf
		for pos := 0; pos < len(testbuf)-4; pos += 4 {
			test1, test2, test3, test4 := awkward[0], awkward[1], awkward[2], awkward[3] // assume decoding uint32 from those
			_, _, _, _ = test1, test2, test3, test4
			awkward = awkward[4:]
		}
	}
}

func BenchmarkSliceMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		awkward := testbuf
		for pos := 0; pos < len(testbuf)-4; pos += 4 {
			test1, test2, test3, test4 := awkward[pos+1], awkward[pos+2], awkward[pos+3], awkward[pos+4] // about same work
			_, _, _, _ = test1, test2, test3, test4
		}
	}
}
