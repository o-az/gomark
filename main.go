package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Headers:           map[string]string{},
			MultiValueHeaders: map[string][]string{},
			Body:              "No body provided",
			IsBase64Encoded:   false,
		}, nil
	}

	var compiledMarkdown = compileMarkdown(request.Body)
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              compiledMarkdown,
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
