package bits

import (
	"encoding/binary"
	"testing"
)

// same buffer, really fast....
// confirms expectation to have benchmarking time match due to inlining of binary.BigEndian* functions

func BenchmarkBinaryUint32(b *testing.B) {
	buf := []byte{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		t := binary.BigEndian.Uint32(buf)
		if t != 16909060 {
			b.Fatal(t) // just for check's sake
		}
	}
}

func BenchmarkCustomUint32(b *testing.B) {
	buf := []byte{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		t := uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
		if t != 16909060 {
			b.Fatal(t) // just for check's sake
		}
	}
}
