package itoa

import "errors"

var BufLenErr = errors.New("Buffer is too short")

func NumToBuf(i uint16, buf []byte) (int, error) {
	digs := 1
	for lim := uint64(10); lim < uint64(i); lim *= 10 {
		digs += 1
	}
	if len(buf) < digs {
		return 0, BufLenErr
	}
	for c := digs - 1; c >= 0 && i > 0; c, i = c-1, i/10 {
		buf[c] = byte(i%10) + '0'
	}
	return digs, nil
}
