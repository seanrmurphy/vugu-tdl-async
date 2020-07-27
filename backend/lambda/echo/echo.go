package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Action  string `json:"action"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// HandleRequest takes a given todo, validates it and posts it to dynamodb
// catching any errors and returning as appropriate; in the case of success, it
// returns the item posted to the dbl
func HandleRequest(e Event) (Response, error) {

	response, _ := json.Marshal(e)
	log.Printf("input event = %v\n", e)
	return Response{
		Status:  200,
		Message: string(response[:]) + "echoed",
	}, nil

	//log.Printf("Successful execution of defaultmessage lambda function\n")
	//return events.APIGatewayProxyResponse{
	////Body:       msg.Content + " (echoed)",
	//Body:       "{\"status\": 200}",
	//StatusCode: 200,
	//}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
