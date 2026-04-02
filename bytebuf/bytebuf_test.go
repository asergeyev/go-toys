package bytebuf

import (
	"bytes"
	"io"
	"testing"
)

// Compares creation and full-read cost of bytes.NewBuffer vs bytes.NewReader.
// bytes.NewReader is read-only and tracks position with an index, while
// bytes.NewBuffer is a read/write type that uses a separate head/tail offset.
// Expectation: NewReader should be slightly cheaper for read-only workloads.

var data = bytes.Repeat([]byte("abcdefghij"), 10000) // 10 MB

func BenchmarkNewBufferCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.NewBuffer(data)
	}
}

func BenchmarkNewReaderCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.NewReader(data)
	}
}

func BenchmarkNewBufferRead(b *testing.B) {
	buf := make([]byte, len(data))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewBuffer(data)
		if _, err := io.ReadFull(r, buf); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewReaderRead(b *testing.B) {
	buf := make([]byte, len(data))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		if _, err := io.ReadFull(r, buf); err != nil {
			b.Fatal(err)
		}
	}
}
