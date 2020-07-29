package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"

	"github.com/seanrmurphy/go-fullstack/backend/model"
	"github.com/seanrmurphy/ws-echo/backend/lambda/types"
	"github.com/seanrmurphy/ws-echo/backend/lambda/util"
)

// Post extracts the Item JSON and writes it to DynamoDB
// Based on https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/dynamodb/create_item.go
func Post(t model.Todo) (model.Todo, error) {
	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// Marshall the Item into a Map DynamoDB can deal with
	av, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return t, err
	}

	// Create Item in table and return
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Todos"),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Error posting to db...error = %v\n", err.Error())
	}
	return t, err
}

// validateTodo performs a couple of basic checks on the todo to ensure it
// contains sensible content before posting it to the database
func validateTodo(t model.Todo) (v model.Todo, e error) {
	v = t
	if t.Title == "" {
		e = errors.New("Invalid Todo Description")
		return
	}
	// limit the status to a specific set...
	nullUuid := uuid.UUID{}
	if t.ID == nullUuid {
		e = errors.New("Invalid UUID")
		return
	}
	return
}

// HandleRequest takes a given todo, validates it and posts it to dynamodb
// catching any errors and returning as appropriate; in the case of success, it
// returns the item posted to the dbl
//func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
func HandleRequest(m types.Message) (types.Response, error) {

	if m.Type != "create-todo" {
		e := util.CreateResponse("NOK", "Handling incorrect message type - ignoring...", "")
		return e, nil
	}

	t := model.Todo{}
	err := json.Unmarshal([]byte(m.Data), &t)
	if err != nil {
		log.Printf("Invalid input - error unmarshalling input%v\n", err.Error())
		e := util.CreateResponse("NOK", "Invalid Todo", "")
		return e, nil
	}

	validTodo, err := validateTodo(t)
	if err != nil {
		log.Printf("Invalid input - should return error %v\n", err.Error())
		e := util.CreateResponse("NOK", "Invalid Todo", "")
		return e, nil
	}

	posted, err := Post(validTodo)

	b, _ := json.Marshal(&posted)
	e := util.CreateResponse("OK", "", string(b))
	return e, nil
}

func main() {
	lambda.Start(HandleRequest)
}
