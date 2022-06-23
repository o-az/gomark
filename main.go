package main

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 400, IsBase64Encoded: false}, nil
	}
	var compiledMarkdown string
	switch request.HTTPMethod {
	case "GET":
		{
			decoded, err := base64.StdEncoding.DecodeString(request.QueryStringParameters["base64"])
			compiledMarkdown = string(decoded)
			if err != nil {
				return events.APIGatewayProxyResponse{StatusCode: 400, IsBase64Encoded: false}, nil
			}
		}
	case "POST":
		{
			compiledMarkdown = compileMarkdown(request.Body)
		}
	default:
		{
			compiledMarkdown = "Invalid method"
		}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: compiledMarkdown}, nil
}

func main() {
	lambda.Start(handleRequest)
}
