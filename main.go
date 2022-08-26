package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"web-telnet-server/pkg/client"

	"github.com/urfave/cli/v2"

	"github.com/gorilla/websocket"
)

func main() {
	address := ""
	endpoint := ""
	app := &cli.App{
		Name:      "wscat",
		Usage:     "a websocket tool.",
		ArgsUsage: " ",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "endpoint",
				Aliases:  []string{"e", "t"},
				Usage:    "WebSocket endpoint, ws:// or wss://",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "listen",
				Aliases:     []string{"l", "p"},
				Usage:       "Listen port to replace stdin, 1337, 127.0.0.1:1337",
				DefaultText: "",
			},
		},
		HideHelp: true,
		Action: func(cCtx *cli.Context) error {
			endpoint = cCtx.String("endpoint")
			address = cCtx.String("listen")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	var conn *net.Conn = nil
	if address != "" {
		port, err := strconv.ParseInt(address, 10, 64)
		if err != nil {
			conn, err = client.TCPListen(address)
		} else {
			conn, err = client.TCPListen(fmt.Sprintf("127.0.0.1:%d", port))
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	wsConn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	wsRecv := make(chan []byte)
	go client.WSReader(wsConn, wsRecv)
	inbond, outbond := make(chan []byte), make(chan []byte)

	if conn == nil {
		go client.StdinHandler(inbond, outbond)
	} else {
		go client.TCPHandler(*conn, inbond, outbond)
	}

	for {
		select {
		case data, ok := <-inbond:
			if !ok {
				return
			}
			err = wsConn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Printf("WebSocket Send err: %v", err)
				return
			}
		case data, ok := <-wsRecv:
			if !ok {
				log.Printf("WebSocket Close")
				return
			}
			outbond <- data
		}
	}
}
