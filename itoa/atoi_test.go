package itoa

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var num = uint16(10000)
var ln = int(5)

func TestItoa(t *testing.T) {
	var err error

	buf := make([]byte, ln)

	for n := uint16(0); n < num; n++ {
		ln, err = NumToBuf(n, buf)
		if answer := []byte(strconv.FormatInt(int64(n), 10)); len(answer) != ln || !bytes.Equal(buf[:ln], answer) {
			t.Errorf("Invalid value returned for %d value: '%s'!='%s'", n, buf[:ln], answer)
		}
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	num = uint16(rand.Uint32())
	str := strconv.FormatUint(uint64(num), 10)
	answer := []byte(str)

	ln, err = NumToBuf(num, buf)
	if err != nil {
		t.Error(err)
	}

	if ln != len(answer) {
		t.Error("Wrong number of bytes returned")
	}
	if !bytes.Equal(buf[:ln], answer) {
		t.Errorf("%s (%d) != %s", buf[:ln], len(buf), answer)
	}
	if _, err = NumToBuf(num, buf[:ln-1]); err == nil {
		t.Error("Shorter buf should've returned error!")
	}
}

func BenchmarkGo(b *testing.B) {
	var str string
	var answer []byte
	for n := 0; n < b.N; n++ {
		str = strconv.FormatUint(uint64(num), 10)
		answer = []byte(str)
		if len(answer) != ln { // same check to make sure variables created
			b.Error("Invalid answer")
		}
	}
}

func BenchmarkNumToBuf(b *testing.B) {
	buf := make([]byte, 100)
	for n := 0; n < b.N; n++ {
		x, _ := NumToBuf(num, buf)
		if x != ln { // same check to make sure variables created
			b.Error("Invalid answer")
		}
	}
}
