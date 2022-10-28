package idgen

import (
	"fmt"
	"testing"
)

func BenchmarkNewStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Str := NewStr()
		_ = Str
	}
}

func BenchmarkFastStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Str := FastStr()
		_ = Str
	}
}

func BenchmarkQuickStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Str := FastStr()
		_ = Str
	}
}

func BenchmarkBinStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Str := BinStr()
		_ = Str
	}
}

func TestNewStr(t *testing.T) {
	var low float64
	for i := 0; i < 1e6; i++ {
		bts := []byte(QuickStr())
		for _, x := range bts {
			if x < b62[31] {
				low++
			}
		}
	}
	fmt.Println(low / 1e6 / 22)
}

func TestFastStr(t *testing.T) {
	fmt.Println(FastStr())
}

func TestQuickStr(t *testing.T) {
	fmt.Println(QuickStr())
}

func TestBinStr(t *testing.T) {
	fmt.Println(BinStr())
}
