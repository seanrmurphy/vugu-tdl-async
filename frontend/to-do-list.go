package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/vugu/vugu"
	"nhooyr.io/websocket"

	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types"
	"github.com/seanrmurphy/vugu-tdl-async/models"
)

func (c *ToDoList) updateItem(t *models.Todo) {

	todoMarshalled, _ := json.Marshal(&t)

	m := types.Message{
		Type: "update-todo",
		Data: string(todoMarshalled[:]),
	}

	messageMarshalled, _ := json.Marshal(&m)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//err = wsjson.Write(ctx, c.conn, v)
	//str := "{\"action\": \"echo\", \"type\": \"t\", \"content\": \"c\"}"
	err := wsConn.Write(ctx, websocket.MessageText, []byte(messageMarshalled))
	if err != nil {
		log.Fatal(err)
	}

}

func (c *ToDoList) postItemToBackend(t models.Todo) {
	//backend := createClient()

	todoMarshalled, _ := json.Marshal(&t)

	m := types.Message{
		Type: "create-todo",
		Data: string(todoMarshalled[:]),
	}

	messageMarshalled, _ := json.Marshal(&m)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//err = wsjson.Write(ctx, c.conn, v)
	//str := "{\"action\": \"echo\", \"type\": \"t\", \"content\": \"c\"}"
	err := wsConn.Write(ctx, websocket.MessageText, []byte(messageMarshalled))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *ToDoList) destroyItemOnBackend(t *models.Todo) {

	m := types.Message{
		Type: "delete-todo",
		Data: t.ID.String(),
	}

	messageMarshalled, _ := json.Marshal(&m)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := wsConn.Write(ctx, websocket.MessageText, []byte(messageMarshalled))
	if err != nil {
		log.Fatal(err)
	}

}

func (c *ToDoList) InitializeTodos(todos []models.Todo) {
	log.Printf("Initializing todo data...")
	c.Todos = make(map[string]models.Todo)
	c.Index = []string{}
	for _, v := range todos {
		idString := v.ID.String()
		c.Todos[idString] = v
		c.Index = append(c.Index, idString)
	}
}

func (c *ToDoList) BeforeBuild() {
}

func (c *ToDoList) getTodoId(s interface{}) (o, id string) {
	slice := strings.SplitN(s.(string), "-", 2)
	o = slice[0]
	id = slice[1]
	return
}

func (c *ToDoList) Done(e vugu.DOMEvent) {
	_, id := c.getTodoId(e.Prop("target", "id"))
	t := c.Todos[id]
	t.Completed = !t.Completed

	c.Todos[id] = t
	go c.updateItem(&t)
}

func (c *ToDoList) Delete(e vugu.DOMEvent) {
	_, id := c.getTodoId(e.Prop("target", "id"))
	log.Printf("Delete - id = %v", id)
	// remove from index
	found := false
	for i, v := range c.Index {
		if v == id {
			switch i {
			case 0:
				found = true
				c.Index = c.Index[i+1:]
				break
			case len(c.Index) - 1:
				found = true
				c.Index = c.Index[:i]
				break
			default:
				found = true
				c.Index = append(c.Index[:i], c.Index[i+1:]...)
				break
			}
		}
	}
	if found == false {
		log.Printf("Unable to remove item from slice")
	}

	// remove from map
	t := c.Todos[id]
	delete(c.Todos, id)
	go c.destroyItemOnBackend(&t)
}

func (c *ToDoList) AddTodo(t models.Todo) {
	if c.Todos == nil {
		c.Todos = make(map[string]models.Todo)
	}
	c.Todos[t.ID.String()] = t
	c.Index = append(c.Index, t.ID.String())
	go c.postItemToBackend(t)
}

func (c *ToDoList) Keypress(e vugu.DOMEvent) {
	keyCode := e.PropFloat64("keyCode")
	// when enter is pressed...
	if keyCode == 13 {
		todoString := e.PropString("target", "value")
		t := models.Todo{ID: strfmt.UUID(uuid.New().String()), Title: &todoString, Completed: false}
		c.AddTodo(t)
	}
}
