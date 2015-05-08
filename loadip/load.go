package loadip

import (
	"errors"
	"net"
)

var AddrError = errors.New("invalid IP address")

// Loadip reads IP address from a string.
func Loadip(ipstr string) (net.IP, error) {
	var (
		ip  uint64
		oct uint32
		b   byte
		num byte
	)

	for _, b = range []byte(ipstr) {
		switch {
		case b == '.':
			num++
			ip = ip<<8 + uint64(oct)
			if ip > 0xffffffff {
				return nil, AddrError
			}
			oct = 0
		case b >= '0' && b <= '9':
			oct = oct*10 + uint32(b-'0')
			if oct > 255 {
				return nil, AddrError
			}
		default:
			return nil, AddrError
		}
	}
	if num != 3 {
		return nil, AddrError
	}
	ip = ip<<8 + uint64(oct)
	if ip > 0xffffffff {
		return nil, AddrError
	}

	return net.IP([]byte{byte(ip >> 24), byte(ip >> 16), byte(ip >> 8), byte(ip)}), nil
}
