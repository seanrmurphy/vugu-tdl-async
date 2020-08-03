package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types"
	"github.com/seanrmurphy/vugu-tdl-async/backend/lambda/util"
	"github.com/seanrmurphy/vugu-tdl-async/models"
)

var tableName string

// GetTodos gets an array of todos and returns them
func GetTodos() (tarray []models.Todo, e error) {

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	proj := expression.NamesList(expression.Name("ID"), expression.Name("Title"), expression.Name("Completed"), expression.Name("CreationDate"))

	//expr, err := expression.NewBuilder().Build()
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		log.Println("Got error building expression:", err.Error())
		e = err
		return
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Println("Query API call failed:", err.Error())
		e = err
		return
	}

	for _, i := range result.Items {
		type simplifiedTodo struct {
			ID           uuid.UUID
			Title        string
			Completed    bool
			CreationDate time.Time
		}

		st := simplifiedTodo{}
		err = dynamodbattribute.UnmarshalMap(i, &st)

		if err != nil {
			log.Println("Got error unmarshalling:", err.Error())
			e = err
			return
		} else {
			t := models.Todo{
				Completed:    st.Completed,
				CreationDate: strfmt.DateTime(st.CreationDate),
				ID:           strfmt.UUID(st.ID.String()),
				Title:        &st.Title,
			}
			tarray = append(tarray, t)
		}
	}

	return
}

// HandleRequest obtains the set of todos from dynamodb
func HandleRequest(m types.Message) (types.Response, error) {

	if m.Type != "list-todos" {
		e := util.CreateResponse("list-todos-response", "NOK", "Handling incorrect message type - ignoring...", "")
		return e, nil
	}

	tableName = os.Getenv("TABLE_NAME")

	// TODO(murp): add some error handling here
	tarray, _ := GetTodos()

	tbody, _ := json.Marshal(tarray)
	return util.CreateResponse("list-todos-response", "OK", "", string(tbody)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
