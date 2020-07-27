package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest takes a given todo, validates it and posts it to dynamodb
// catching any errors and returning as appropriate; in the case of success, it
// returns the item posted to the dbl
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Executing onconnect lambda function\n")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
