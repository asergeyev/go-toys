package loadip

import (
	"fmt"
	"net"
	"testing"
	"time"
)

var testIPs = []string{
	"0.0.0.0",
	"127.0.0.1",
	"255.255.255.255",

	// Errors
	"1.2.3",
	"2.3.4.256",
	"2.3.4.-1",
	"2.3.4.5.6",
}

func TestLoadip(t *testing.T) {
	for _, ip := range testIPs {
		expect := net.ParseIP(ip)
		reslt, _ := Loadip(ip)
		if expect == nil && reslt != nil {
			t.Errorf("IP %s resolved to %s but net.ParseIP gives nil", ip, reslt.String())
		} else if expect != nil && reslt == nil {
			t.Errorf("IP %s resolved to nil but net.ParseIP gives %s", ip, expect.String())
		} else if !net.ParseIP(ip).Equal(reslt) {
			t.Errorf("IP %s resolved to %s but net.ParseIP gives %s", ip, reslt.String(), expect.String())
		}
	}
}

var Sec = time.Now().Unix()
var BenchIP = fmt.Sprintf("%d.%d.%d.%d", byte(Sec>>24), byte(Sec>>16), byte(Sec>>8), byte(Sec))

func BenchmarkLoadIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, _ := Loadip(BenchIP)
		if s == nil {
			b.Fatal("Something is not right in BenchmarkLoadIP")
		}
	}
}

func BenchmarkParseIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := net.ParseIP(BenchIP)
		if s == nil {
			b.Fatal("Something is not right in BenchmarkParseIP")
		}
	}
}
