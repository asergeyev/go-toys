package reuse

import (
	"net"

	"testing"
)

func TestServer(t *testing.T) {
	var err error

	port := 5555
	workers := make([]*Service, 5)

	for i := 0; i < len(workers); i++ {
		workers[i], err = NewService(port)
		if err != nil {
			t.Fatalf("Failed to create clone #%d %s", i, err)
		}
	}

	// success, now prove it's not possible to Listen on same..

	addr := workers[len(workers)-1].l.LocalAddr() // this one is definitely cloned

	_, err = net.ListenUDP("udp", addr.(*net.UDPAddr))
	if err == nil {
		t.Fatal("Expectations failed, Listen worked")
	}
}
