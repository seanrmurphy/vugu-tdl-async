package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types"
	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/util"
	"github.com/seanrmurphy/vugu-tdl-async/models"
)

var tableName string

// UpdateTodo updates a given todo in the dynamodb database
func UpdateTodo(t models.Todo) (models.Todo, error) {

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	u, err := uuid.Parse(t.ID.String())
	if err != nil {
		fmt.Println(err.Error())
		return t, err
	}

	d := time.Time(t.CreationDate)

	uuidBinary, _ := u.MarshalBinary()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":t": {
				S: aws.String(*t.Title),
			},
			":d": {
				S: aws.String(d.Format(time.RFC3339Nano)),
			},
			":c": {
				BOOL: aws.Bool(t.Completed),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				B: uuidBinary,
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Title = :t, Completed = :c, CreationDate = :d"),
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return t, err
	}

	return t, nil
}

// HandleRequest performs a simple check to ensure that the id provided is valid
// and calls the update function; it returns the updated record jsonified
func HandleRequest(m types.Message) (types.Response, error) {

	if m.Type != "update-todo" {
		e := util.CreateResponse("update-todo-response", "NOK", "Handling incorrect message type - ignoring...", "")
		return e, nil
	}

	tableName = os.Getenv("TABLE_NAME")

	//id := req.PathParameters["todoid"]
	//if id == "" {
	//return events.APIGatewayProxyResponse{
	//StatusCode: http.StatusInternalServerError,
	//Body:       "No ID provided",
	//}, nil
	//}

	t := models.Todo{}
	_ = json.Unmarshal([]byte(m.Data), &t)
	var err error
	//t.ID, err = uuid.Parse(id)
	log.Printf("Received Todo: %s\n", t)

	returnedTodo := models.Todo{}
	returnedTodo, err = UpdateTodo(t)
	if err != nil {
		e := util.CreateResponse("update-todo-response", "NOK", "Error updating todo in DB...", "")
		return e, nil
	}

	tbody, _ := json.Marshal(returnedTodo)
	return util.CreateResponse("update-todo-response", "OK", "", string(tbody)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
