package atoi

// #include <stdlib.h>
// int atoi2(void *buf){return atoi((char *)buf); }
import "C"
import "unsafe"

func CAtoi(buf []byte) (int, error) {
	// error is always nil
	return int(C.atoi2(unsafe.Pointer(&buf[0]))), nil
}
