package main

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(
	ctx context.Context,
	request events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {
	var compiledMarkdown string

	switch request.RequestContext.HTTP.Method {

	case "GET":
		text := request.QueryStringParameters["text"]
		if len(text) == 0 {
			return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "Need 'text' query param"}, nil
		}
		decoded, err := base64.StdEncoding.DecodeString(text)
		compiledMarkdown = string(decoded)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "Error decoding text query param"}, nil
		}
		return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: compiledMarkdown}, nil

	case "POST":
		if len(request.Body) == 0 {
			return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "Need body in 'POST' request"}, nil
		}
		compiledMarkdown = compileMarkdown(request.Body)
		return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: compiledMarkdown}, nil

	default:
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Must be GET or POST. Need 'text' query param in GET or body in POST",
		}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
