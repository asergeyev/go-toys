package floats

import (
	"testing"
)

// this illustrates that cost of two int multiplications is as high as float operation.
// this also brings up the fact that native length floats are better than not-native (float32
// test would be slower on amd64)

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := int(float32(i) * 1.75)
		if x == -1 {
			b.Error("fake error")
		}
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := int(float64(i) * 1.75)
		if x == -1 {
			b.Error("fake error")
		}
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := i * 7 / 4
		if x == -1 {
			b.Error("fake error")
		}
	}
}
