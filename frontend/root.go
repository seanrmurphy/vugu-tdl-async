package main

import (
	"log"

	"github.com/seanrmurphy/vugu-tdl-swagger/swagger/models"
)

var websocketService = "wss://x6ajz85dla.execute-api.eu-west-1.amazonaws.com/Prod"

func (c *Root) initializeTodos(t []models.Todo) {
	if c.Body == nil {
		log.Printf("Unable to initialize todos\n")
		return
	}
	c.events.Lock()
	c.Body.(*ToDoList).InitializeTodos(t)
	c.events.UnlockRender()
}
