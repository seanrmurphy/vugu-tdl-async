// +build wasm

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"flag"

	"github.com/seanrmurphy/ws-echo/backend/lambda/types"
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/domrender"
	"nhooyr.io/websocket"
)

var wsConn *websocket.Conn

func main() {

	mountPoint := flag.String("mount-point", "#vugu_mount_point", "The query selector for the mount point for the root component, if it is not a full HTML component")
	flag.Parse()

	fmt.Printf("Entering main(), -mount-point=%q\n", *mountPoint)
	defer fmt.Printf("Exiting main()\n")

	buildEnv, err := vugu.NewBuildEnv()
	if err != nil {
		panic(err)
	}

	renderer, err := domrender.New(*mountPoint)
	if err != nil {
		panic(err)
	}
	defer renderer.Release()

	rootBuilder := vuguSetup(buildEnv, renderer.EventEnv())
	//rootBuilder := &Root{}

	initWSConnection()

	for ok := true; ok; ok = renderer.EventWait() {

		buildResults := buildEnv.RunBuild(rootBuilder)

		err = renderer.Render(buildResults)
		if err != nil {
			panic(err)
		}
	}

}

func initWSConnection() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var err error
	wsConn, _, err = websocket.Dial(ctx, websocketService, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Opened websocket connection...\n")

	// request Todos
	m := types.Message{
		Type: "list-todos",
		Data: "",
	}

	messageMarshalled, _ := json.Marshal(&m)

	//err = wsjson.Write(ctx, c.conn, v)
	//str := "{\"action\": \"echo\", \"type\": \"t\", \"content\": \"c\"}"
	err = wsConn.Write(ctx, websocket.MessageText, []byte(messageMarshalled))
	if err != nil {
		log.Fatal(err)
	}

	go receiver()
}

// this function simply wait for a response from a websocket...
func receiver() {
	for {
		ctx := context.TODO()
		_, msgMarshalled, err := wsConn.Read(ctx)
		if err != nil {
			log.Printf("Error reading message %v\n", err)
		} else {
			msg := types.Response{}
			_ = json.Unmarshal(msgMarshalled, &msg)
			log.Printf("Message successfully received - type = %v, data = %v\n", msg.Type, msg.Data)
		}
	}
}
