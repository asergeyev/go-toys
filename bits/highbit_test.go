package bits

import (
	"math/rand"

	"testing"
)

func highbit(c uint32) (n uint8) {
	if c >= 65536 {
		c >>= 16
		n += 16
	}
	if c >= 256 {
		c >>= 8
		n += 8
	}
	if c >= 16 {
		c >>= 4
		n += 4
	}
	if c >= 4 {
		if c >= 8 {
			return n + 4
		} else {
			return n + 3
		}
	} else {
		if c >= 2 {
			return n + 2
		} else if c != 0 {
			return n + 1
		}
	}
	return n
}

func TestHighBit(t *testing.T) {
	if highbit(0) != 0 {
		t.Error("invalid for 0")
	}
	if highbit(1) != 1 {
		t.Error("invalid for 1")
	}
	var (
		n    uint32
		resp uint8 = 1
	)
	for n = 2; n >= 2; n *= 2 { // eventually will go to 0
		resp++
		if x := highbit(n); x != resp {
			t.Error("invalid for", n, "got", x, "expected", resp)
		}
		if x := highbit(n + 1); x != resp { // 1 should not change result
			t.Error("invalid for", n+1, "got", x, "expected", resp)
		}
		if x := highbit(n - 1); x != resp-1 { // -1 should decrease highest bit
			t.Error("invalid for", n-1, "got", x, "expected", resp-1)
		}
	}
}

func BenchmarkHighbitZero(b *testing.B) {
	var res uint8
	for i := 0; i < b.N; i++ {
		res += highbit(0)
	}
}

var testcase []uint32

func init() {
	testcase = make([]uint32, 256) // worst case is 0?
	for i := range testcase {
		testcase[i] = rand.Uint32()
	}
}

func BenchmarkHighbit(b *testing.B) {
	var res uint8
	for i := 0; i < b.N; i++ { // this one is expected to be slower, since it has to lookup data in array
		res += highbit(testcase[byte(i)])
	}
}

func BenchmarkHighbitsingle(b *testing.B) {
	c := rand.Uint32()

	var res uint8
	for i := 0; i < b.N; i++ {
		res += highbit(c)
	}
}
