package mapaccess

// this one is to estimate how much more expensive access to values in map[string]...
// vs access to map[int] ones...

import (
	"testing"
)

func BenchmarkMapAccess2(b *testing.B) {
	var test = make(map[string]int)
	for a := ' '; a <= 'z'; a++ {
		test[string(a)] = 1
		test["x"+string(a)] = 1
	}

	m := '0'
	var ok bool
	for i := 0; i < b.N; i++ {
		if _, ok = test["x"+string(m)]; ok {
		}
		if m == 255 {
			m = ' '
		} else {
			m++
		}
	}
}

func BenchmarkMapAccess1(b *testing.B) {
	var test = make(map[string]int)
	for a := ' '; a <= 'z'; a++ {
		test[string(a)] = 1
		test["x"+string(a)] = 1
	}

	m := '0'
	var s int
	for i := 0; i < b.N; i++ {
		if s = test["x"+string(m)]; s > 0 {
		}
		if m == 255 {
			m = ' '
		} else {
			m++
		}
	}
}

func BenchmarkMapAccessInt(b *testing.B) {
	var test = make(map[int]int)
	for a := ' '; a <= 'z'; a++ {
		test[int(a)] = 1
		test[1000+int(a)] = 1
	}

	m := '0'
	var ok bool
	for i := 0; i < b.N; i++ {
		if _, ok = test[1000+int(m)]; ok {
		}
		if m == 255 {
			m = ' '
		} else {
			m++
		}
	}
}
