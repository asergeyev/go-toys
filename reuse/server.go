package reuse

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

/*
#include <sys/socket.h>
*/
import "C"

type Service struct {
	l net.Conn
}

func NewService(port int) (*Service, error) {
	// Make socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return nil, err
	}
	syscall.CloseOnExec(fd)

	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		return nil, err
	}
	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, C.SO_REUSEPORT, 1); err != nil {
		return nil, err
	}
	err = syscall.SetNonblock(fd, true)
	if err != nil {
		return nil, err
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: port, Addr: [4]byte{127, 0, 0, 1}})
	if err != nil {
		return nil, err
	}

	conn, err := net.FileConn(os.NewFile(uintptr(fd), fmt.Sprintf("udp:127.0.0.1:%d", port)))
	if err != nil {
		return nil, err
	}

	return &Service{conn}, nil
}
