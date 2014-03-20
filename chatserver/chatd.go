package main

// chatd - listens on port 7037 and transmits received '\n' 
// terminated lines to all other connected clients.

// This work demonstrates possible approach to demultiplexing.


import (
	"bufio"
	"io"
	"log"
	"net"
)

type Client interface {
	io.ReadWriter
	Id() string
}

type Event struct {
	clid    string
	message []byte
}

type Console struct {
	cn  chan Client
	fin chan string
}

func (c *Console) Stream(cl Client, out chan Event) {
	myid := cl.Id()
	buf := bufio.NewReader(cl)
	for {
		line, err := buf.ReadSlice('\n')
		if line != nil {
			out <- Event{myid, line}
		}
		if err != nil {
			break
		}
	}
	c.fin <- myid
}

func (c *Console) Send(cl Client, msg []byte) {
	_, e := cl.Write(msg)
	if e != nil {
		c.fin <- cl.Id()
	}
}

func (c *Console) Multiplex() {
	messages := make(chan Event)
	conn := make(map[string]Client)
	for {
		select {
		case cli := <-c.cn:
			conn[cli.Id()] = cli
			go c.Stream(cli, messages)
		case fin := <-c.fin:
			delete(conn, fin)
		case e := <-messages:
			for id, cli := range conn {
				if id != e.clid {
					go c.Send(cli, e.message)
				}
			}
		}
	}
}

type chatClient struct {
	connection net.Conn
}

func NewChatClient(conn net.Conn) *chatClient {
	return &chatClient{conn}
}

func (cl *chatClient) Write(b []byte) (int, error) {
	return cl.connection.Write(b)
}

func (cl *chatClient) Id() string {
	return cl.connection.RemoteAddr().String()
}

func (cl *chatClient) Read(b []byte) (int, error) {
	return cl.connection.Read(b)
}

func main() {
	console := &Console{make(chan Client, 1), make(chan string, 1)}

	go console.Multiplex()

	ln, err := net.Listen("tcp", ":7037")
	if err != nil {
		log.Fatal("Unable to listen on :7037")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Unable to listen on :7037")
			continue
		}
		console.cn <- NewChatClient(conn)
	}
}
