package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type request events.APIGatewayProxyRequest
type response events.APIGatewayProxyResponse

func handler(req request) (response, error) {
	return response{Body: jsonToUpper(req.Body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
