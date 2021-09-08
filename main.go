package main

import "github.com/aws/aws-lambda-go/lambda"
import "github.com/aws/aws-lambda-go/events"

func handleRequest () (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body: "\"Hello from Go!\"",
    }, nil
}

func main() {
    lambda.Start(handleRequest)
}
