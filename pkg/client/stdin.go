package client

import (
	"bufio"
	"fmt"
	"os"
)

func StdinHandler(recv, send chan []byte) {
	go StdinSender(send)
	StdinReader(recv)
}

func StdinReader(recv chan []byte) {
	reader := bufio.NewReader(os.Stdin)
	for {
		buf := make([]byte, BufferSize)
		n, err := reader.Read(buf)

		if err != nil {
			close(recv)
			return
		}
		recv <- buf[:n]
	}
}

func StdinSender(send chan []byte) {
	for data := range send {
		fmt.Print(string(data))
	}
}
