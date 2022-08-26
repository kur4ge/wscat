package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	BufferSize = 2048
)

func TCPListen(address string) (*net.Conn, error) {
	fmt.Printf("Listen %s...\n", address)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	conn, err := listen.Accept()
	fmt.Printf("Received connection from %s\n", conn.RemoteAddr())
	return &conn, err
}

func TCPHandler(conn net.Conn, inbond, outbond chan []byte) {
	go TCPSender(conn, outbond)
	TCPReader(conn, inbond)
}

func TCPReader(conn net.Conn, inbond chan []byte) {
	reader := bufio.NewReader(conn)
	for {
		buf := make([]byte, BufferSize)
		n, err := reader.Read(buf)
		if err != nil {
			log.Printf("TCP Close")
			return
		}
		fmt.Printf("%s", string(buf[:n]))
		inbond <- buf[:n]
	}
}

func TCPSender(conn net.Conn, outbond chan []byte) {
	for data := range outbond {
		fmt.Printf("%s", string(data))
		_, err := conn.Write(data)
		if err != nil {
			log.Printf("TCP Send err: %v", err)
			return
		}
	}
}
