package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// HandleRequest takes a given todo, validates it and posts it to dynamodb
// catching any errors and returning as appropriate; in the case of success, it
// returns the item posted to the dbl
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	msg := Message{}
	log.Printf("req.Body = %v\n", req.Body)
	if err := json.Unmarshal([]byte(req.Body), &msg); err != nil {
		log.Printf("Executing defaultmessage lambda function\n")
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error parsing message",
		}, nil
	}

	log.Printf("Successful execution of defaultmessage lambda function\n")
	return events.APIGatewayProxyResponse{
		//Body:       msg.Content + " (echoed)",
		Body:       "{\"status\": 200}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
