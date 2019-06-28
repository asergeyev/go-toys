package wrap

import (
	"testing"
	"github.com/pkg/errors"
)

func BenchmarkWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := errors.Wrap(nil, "test")
		if ts := errors.Wrap(err, "test"); ts != nil {
			panic("not here")
		}
	}
}
func BenchmarkNoWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := errors.Wrap(nil, "test")
		if err != nil {
			panic("not here")
		}
	}
}
