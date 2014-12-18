package atoi

import (
	"strconv"
	"testing"
)

var buf = []byte("123456")

func TestAtoi(t *testing.T) {
	i1, _ := strconv.Atoi(string(buf))
	i2, _ := CAtoi(buf)
	if i1 != i2 {
		t.Errorf("Answer from CAtoi is different than from Go version (%d)", i2)
	}
}

func BenchmarkGo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.Atoi(string(buf))
	}
}

func BenchmarkC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CAtoi(buf)
	}
}
