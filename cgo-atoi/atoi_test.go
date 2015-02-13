package atoi

import (
	"strconv"
	"testing"
)

// note this only works on null-terminated buffer and length 6
// since it's a very simple toy
var buf = append([]byte("12345"), 0)

func TestAtoi(t *testing.T) {
	i1, _ := strconv.Atoi(string(buf[:5]))
	i2, _ := CAtoi(buf)
	if i1 != i2 {
		t.Errorf("Answer from CAtoi is different than from Go version (%d)", i2)
	}
	i3, _ := Direct(buf[:5])
	if uint16(i1) != i3 {
		t.Errorf("Answer from ParseBytes is different than from Go version (%d)", i3)
	}
	i4, _ := DirectShift(buf[:5])
	if uint16(i1) != i4 {
		t.Errorf("Answer from ParseBytes is different than from Go shift version (%d)", i3)
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

func BenchmarkDirect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Direct(buf)
	}
}

func BenchmarkDirectShift(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DirectShift(buf)
	}
}
