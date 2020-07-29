package main

import (
	"github.com/vugu/vugu"
)

type Message struct {
	Action  string `json:"action"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

var websocketService = "wss://x6ajz85dla.execute-api.eu-west-1.amazonaws.com/Prod"

//func (c *Root) init() {

//ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
//defer cancel()

//var err error
//c.conn, _, err = websocket.Dial(ctx, websocketService, nil)
//if err != nil {
//log.Fatal(err)
//}
//log.Printf("Opened websocket connection...\n")

//go c.receiver()
//}

// this function simply wait for a response from a websocket...
//func (c *Root) receiver() {
//for {
//ctx := context.TODO()
//_, msg, err := c.conn.Read(ctx)
//if err != nil {
//log.Printf("Error reading message %v\n", err)
//} else {
//log.Printf("Message successfully received %v\n", string(msg))
//}
//}
//}

func (c *Root) Keypress(e vugu.DOMEvent) {
	//keyCode := e.PropFloat64("keyCode")
	//// when enter is pressed...
	//if keyCode == 13 {
	//v := e.PropString("target", "value")

	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()

	////err = wsjson.Write(ctx, c.conn, v)
	////str := "{\"action\": \"echo\", \"type\": \"t\", \"content\": \"c\"}"
	//err := c.conn.Write(ctx, websocket.MessageText, []byte(v))
	//if err != nil {
	//log.Fatal(err)
	//}

	//}
}
