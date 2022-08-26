package client

import (
	"log"

	"github.com/gorilla/websocket"
)

func WSReader(conn *websocket.Conn, recv chan []byte) {
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			close(recv)
			return
		}
		if len(data) == 0 {
			continue
		}
		switch messageType {
		case websocket.TextMessage: // Do not process data of type TextMessage, output as error message
			log.Println(string(data))
			continue
		}
		recv <- data
	}
}
