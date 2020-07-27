package main

import (
	"context"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var server = "wss://localhost:5000"

func createConnection() ws.Dialer {
	d := ws.Dialer{}
	c := context.TODO()
	_, _, _, _ = d.Dial(c, server)
	return d
}

func main() {
	// open connection to server
	c := createConnection()

	for {
		// send a message
		m, err := wsutil.ReadMessage()
		// wait for a reply
	}

}
