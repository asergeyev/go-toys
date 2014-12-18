package atoi

import (
	"strconv"
	"testing"
)

// note this only works on null-terminated buffer and length 6
// since it's a very simple toy
var buf = append([]byte("123456"), 0)

func TestAtoi(t *testing.T) {
	i1, _ := strconv.Atoi(string(buf[:6]))
	i2, _ := CAtoi(buf)
	if i1 != i2 {
		t.Errorf("Answer from CAtoi is different than from Go version (%d)", i2)
	}
}

func BenchmarkGo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.Atoi(string(buf[:6]))
	}
}

func BenchmarkC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CAtoi(buf)
	}
}
