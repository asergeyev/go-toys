package itoa

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var num = 1000000

func TestNumToBuf(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))

	buf := make([]byte, 100)
	test := make([]byte, 100)

	for n := 0; n < num; n += int(rand.Int63n(10000)) {
		ln, err := NumToBuf(uint64(n), buf)
		if err != nil {
			t.Errorf("Unable to convert %d, %s", n, err)
		}
		test = strconv.AppendUint(test[:0], uint64(n), 10)
		if !bytes.Equal(buf[:ln], test) {
			t.Errorf("%d gives error: %s != %s", n, string(buf[:ln]), test)
			break
		}
	}
}

func BenchmarkItoa(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = strconv.Itoa(n)
		_ = str
	}
}

func BenchmarkFormat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = strconv.FormatInt(int64(n), 10)
		_ = str
	}
}

func BenchmarkToBuf(b *testing.B) {
	buf := make([]byte, 100)
	for n := 0; n < b.N; n++ {
		x, _ := NumToBuf(uint64(num), buf)
		_ = x
	}
}

func BenchmarkSprintf(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%d", n)
		_ = str
	}
}

func BenchmarkSprint(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprint(num)
		_ = str
	}
}

func BenchmarkAppend(b *testing.B) {
	str := make([]byte, 0, 0)
	for n := 0; n < b.N; n++ {
		str = strconv.AppendInt(str[:0], int64(n), 10)
		_ = str
	}
}
