package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/google/uuid"

	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types"
	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/util"
)

var tableName string

// DeleteTodo deletes a todo specified with a uuid. In the case that this does
// not exist an error is generated.
func DeleteTodo(id uuid.UUID) (err error) {

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	uuidBinary, _ := id.MarshalBinary()
	var resp *dynamodb.DeleteItemOutput
	resp, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				B: uuidBinary,
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	})

	// unlikely that an error occurs...
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("Unable to delete item with given ID")
		return
	}

	// check if the return value returned a sensible value
	if _, ok := resp.Attributes["ID"]; !ok {
		// not botherig to confirm ID is correct here; this should be done within
		// dynamodb
		err = errors.New("Item with given ID not found")
	}
	return
}

// HandleRequest performs some basic validation on the input id, if valid sends
// to the delete function and generates a return JSON string which is human readabled
func HandleRequest(m types.Message) (types.Response, error) {

	if m.Type != "delete-todo" {
		e := util.CreateResponse("delete-todo-response", "NOK", "Handling incorrect message type - ignoring...", "")
		return e, nil
	}

	tableName = os.Getenv("TABLE_NAME")

	idString := m.Data
	if idString == "" {
		e := util.CreateResponse("delete-todo-response", "NOK", "No valid ID provided", "")
		return e, nil
	}

	id, _ := uuid.Parse(idString)
	err := DeleteTodo(id)

	if err != nil {
		e := util.CreateResponse("delete-todo-response", "NOK", "No object with given ID found", "")
		return e, nil
	}

	//return fmt.Sprintf("Hello %s!", name.Name), nil
	e := util.CreateResponse("delete-todo-response", "OK", "Record deleted", "")
	return e, nil
}

// main starts the lambda function
func main() {
	lambda.Start(HandleRequest)
}
