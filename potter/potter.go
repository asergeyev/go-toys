package potter

import (
	"bytes"
)

type Pot struct {
	buffer [][]byte
	refs   map[uint32][]byte
}

const BUFSIZE = uint32(1024 * 1024)

func NewPot(seats int) *Pot {
	return &Pot{
		buffer: [][]byte{make([]byte, 0, BUFSIZE)},
		refs:   make(map[uint32][]byte, seats),
	}
}

func hashval(buf []byte) uint32 {
	hash := uint32(5381)
	for _, b := range buf {
		hash = (hash + (hash << 5)) ^ uint32(b)
	}
	return hash
}

func (p *Pot) Find(buf []byte) (h uint32, _ bool) {
	h = hashval(buf)
	for p.refs[h] != nil {
		if bytes.Equal(p.refs[h], buf) {
			return h, true
		}
		h++
	}
	return h, false
}

func (p *Pot) Encode(buf []byte) uint32 {
	h, ok := p.Find(buf)
	if ok {
		return h
	}
	current := len(p.buffer) - 1
	start := len(p.buffer[current])
	if start+len(buf) > cap(p.buffer[current]) {
		current++
		start = 0
		p.buffer = append(p.buffer, make([]byte, 0, BUFSIZE))
	}
	p.buffer[current] = append(p.buffer[current], buf...)
	p.refs[h] = p.buffer[current][start : start+len(buf)]
	return h

}

func (p *Pot) Decode(n uint32) []byte {
	return p.refs[n]
}
