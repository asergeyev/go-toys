// potter - small personal toy project, abstract map[[]byte]uint64 implementation... Feel free to use uint64 for storing unsafe pointers.
// Imagined for relatively large amounts of byte values and no expectation of iterating over them.
package potter

import (
	"bytes"
	"sync"
)

// Pot is a map[uint32]uint32 implementation
type Pot struct {
	buffers [][]byte
	hash    []uint32
	pos     []uint64
	ln      []byte
	points  uint32
	*sync.RWMutex
}

// BufferSize controls amount of bytes per slice of keys buffer. You can change it any time, it gets used by NewPot and by
// Encode when new slice is necessary to add.
var BufferSize uint32 = 5 * 1024 * 1024 // do not overflow it!

// NewPot reserves space for hash elements.
func NewPot(seats uint32) *Pot {
	spaces := int(float64(seats) / 3.0 * 4)
	return &Pot{
		buffers: [][]byte{make([]byte, 0, BufferSize)},
		hash:    make([]uint32, spaces),
		pos:     make([]uint64, spaces),
		ln:      make([]byte, spaces),
		points:  seats,
		RWMutex: &sync.RWMutex{},
	}
}

// Decode digs out value by given int representation.
func (p *Pot) Decode(x uint32) []byte {
	p.RLock()
	defer p.RUnlock()
	if p.pos[x] == 0 {
		// this is not here at all
		return nil
	}
	start := uint32(p.pos[x])
	current := p.buffers[uint32(p.pos[x]>>32)-1]
	return current[start : start+uint32(p.ln[x])]
}

func hashval(buf []byte) uint32 {
	// basuc cdb hash
	hash := uint32(5381)
	for _, b := range buf {
		hash = (hash + (hash << 5)) ^ uint32(b)
	}
	return hash
}

// Encode checks if value already seen and returns it's position; if value is not seen it's added to buffer and new position returned.
func (p *Pot) Encode(buf []byte) uint32 {
	if len(buf) > 255 {
		panic("long buffer could not be encoded")
	}
	hash := hashval(buf)

	p.RLock()
	for i := 0; i < len(p.ln); i++ {
		x := (hash + uint32(i)) % uint32(len(p.ln))
		if p.pos[x] == 0 {
			p.RUnlock()
			p.Lock()
			if p.points == 0 {
				p.resize()
			}
			current := p.buffers[len(p.buffers)-1]
			if len(current)+len(buf) > cap(current) {
				current = make([]byte, 0, BufferSize)
				p.buffers = append(p.buffers, current)
			}
			p.ln[x] = byte(len(buf))
			p.hash[x] = hash
			p.pos[x] = uint64(len(p.buffers))<<32 + uint64(len(current))
			current = append(current, buf...)
			p.buffers[len(p.buffers)-1] = current
			p.points -= 1
			p.Unlock()
			return uint32(x)
		} else if p.hash[x] == hash && byte(len(buf)) == p.ln[x] {
			start := uint32(p.pos[x])
			current := p.buffers[uint32(p.pos[x]>>32)-1]
			if start+uint32(len(buf)) <= uint32(len(current)) {
				val := current[start : start+uint32(len(buf))]
				if bytes.Equal(val, buf) {
					p.RUnlock()
					return x
				}
			}
		}
	}
	panic("pot is full")
}

// SendAllKeys streams all taken hashtab positions to channel
func (p *Pot) SendAllKeys(x chan int) {
	p.RLock()
	defer p.RUnlock()
	for slot, pos := range p.pos {
		if pos > 0 {
			x <- slot
		}
	}
	close(x)
}

func (p *Pot) resize() {
	// no locking here

	// generate new arrays for ((currentsize+1) * 1.6)
	addsize := float64(len(p.ln)+1) * 0.6
	p.points += uint32(addsize*.75 + .5)

	newlen := uint32(len(p.ln) + int(addsize+.5))
	newln, newhash, newpos := make([]byte, newlen), make([]uint32, newlen), make([]uint64, newlen)

HASH:
	// replay the hashes and insert new values
	for x, hash := range p.hash {
		for i := 0; i < len(p.ln); i++ {
			newx := (hash + uint32(i)) % newlen
			if newpos[newx] == 0 {
				newpos[newx] = p.pos[x]
				newln[newx] = p.ln[x]
				newhash[newx] = hash
				continue HASH
			}
		}
		panic("could not resize, all spots are taken!")
	}

	// swap old arrays to new ones
	p.ln, p.hash, p.pos = newln, newhash, newpos
}
