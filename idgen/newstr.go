package idgen

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"math/bits"
)

func NewStr() string {
	var bytes [17]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		panic(err)
	}

	var i big.Int
	i.SetBytes(bytes[:])

	u := i.Text(62)
	if l := len(u); l < 22 {
		return fmt.Sprintf("%022s", u)
	} else {
		return u[l-22:]
	}
}

const (
	b62  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	blen = byte(len(b62))
)

// need to take 6 bit at a time
// 111101 is max in base62 so if 11111? is pulled, then
// break is called and we shift by one bit forward

func QuickStr() string {
	slice := make([]byte, 22)
	if _, err := io.ReadAtLeast(rand.Reader, slice, 22); err != nil {
		panic(err)
	}
	for pos, b := range slice {
		slice[pos] = b62[b%62]
	}
	return string(slice)
}

func FastStr() string {
	var bytes [17]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		panic(err)
	}

	ret := [22]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	var next, bpos, bit byte
	for pos := byte(0); int(pos) < len(ret); pos++ {
		bpos, bit = (pos*6)/8, (pos*6)%8
		if bit < 2 {
			next += (bytes[bpos] >> (2 - bit))
		} else {
			// otherwise it's few bits from this byte and few from the next
			// next = ((bytes[bpos] & 31) << 1) | (bytes[bpos+1]>>7) // bit=2
			// next = ((bytes[bpos] & 15) << 2) | (bytes[bpos+1]>>6) // bit=3
			// next = ((bytes[bpos] & 7) << 3) | (bytes[bpos+1]>>5) //  bit=4

			// first AND value is 1<<(7-bit)-1, save it as temp value
			// fitst SHIFT value is bit-1
			// last SHIFT is 9-bit
			next += ((bytes[bpos] & ((1 << (7 - bit)) - 1)) << (bit - 1)) | (bytes[bpos+1] >> (9 - bit))
		}
		ret[pos] = b62[next%32]
	}

	return string(ret[:])
}

func BinStr() []byte {
	buf := make([]byte, 24)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	out := buf[:22]
	q0 := binary.LittleEndian.Uint64(buf)
	q1 := binary.LittleEndian.Uint64(buf)
	q2 := binary.LittleEndian.Uint64(buf)
	var r0, r1, r2 uint64
	for i := 0; i < 22; i += 3 {
		q0, r0 = bits.Div64(0, q0, 62)
		q1, r1 = bits.Div64(0, q1, 62)
		q2, r2 = bits.Div64(0, q2, 62)
		out[i] = b62[r0] // 0,3,6,21
		if i < 21 {
			out[i+1] = b62[r1]
			out[i+2] = b62[r2] //
		}
	}

	return out
}
