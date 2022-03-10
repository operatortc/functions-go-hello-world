package main

import "github.com/aws/aws-lambda-go/lambda"
import "github.com/aws/aws-lambda-go/events"

import "encoding/json"
import "context"

func handleRequest (ctx context.Context, event json.RawMessage) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body: "\"Hello from Go!\"",
    }, nil
}

func main() {
    lambda.Start(handleRequest)
}
