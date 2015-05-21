package potter

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewPot(t *testing.T) {
	x := NewPot(10000)
	for i := 0; i < 65535; i++ {
		core := []byte{byte(i), byte(i >> 8), byte(rand.Int31()), byte(rand.Int31())}
		pt := x.Encode(core)
		ret := x.Decode(pt)
		if len(ret) != 4 || ret[0] != core[0] || ret[1] != core[1] {
			t.Errorf("Invald potting %v != %v", ret, core)
		}
	}
}

func BenchmarkPotting(b *testing.B) {
	x := NewPot(b.N)
	for i := 0; i < b.N; i++ {
		x.Encode([]byte(fmt.Sprint(i)))
	}
}

func BenchmarkPottingLookupUnknown(b *testing.B) {
	test := make([][]byte, b.N)
	x := NewPot(b.N)
	for i := 0; i < b.N; i++ {
		test[i] = []byte(fmt.Sprint("abcdef", i))
		x.Encode([]byte(fmt.Sprint(i)))
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		x.Encode(test[j])
	}
}

func BenchmarkPottingLookupKnown(b *testing.B) {
	test := make([][]byte, b.N)
	x := NewPot(b.N)
	for i := 0; i < b.N; i++ {
		test[i] = []byte(fmt.Sprint(i))
		x.Encode([]byte(fmt.Sprint(i)))
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		x.Encode(test[j])
	}
}
