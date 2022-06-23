package main

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var compiledMarkdown string

	if request.HTTPMethod == "GET" {
		text := request.QueryStringParameters["text"]
		if len(text) == 0 {
			return events.APIGatewayProxyResponse{
				StatusCode:        400,
				Headers:           map[string]string{},
				MultiValueHeaders: map[string][]string{},
				Body:              "Need 'text' query param",
				IsBase64Encoded:   false,
			}, nil
		}
		decoded, err := base64.StdEncoding.DecodeString(text)
		compiledMarkdown = string(decoded)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode:        400,
				Headers:           map[string]string{},
				MultiValueHeaders: map[string][]string{},
				Body:              "Error decoding text query param",
				IsBase64Encoded:   false,
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode:        200,
			Headers:           map[string]string{},
			MultiValueHeaders: map[string][]string{},
			Body:              compiledMarkdown,
			IsBase64Encoded:   false,
		}, nil
	}

	if request.HTTPMethod == "POST" {
		if len(request.Body) == 0 {
			return events.APIGatewayProxyResponse{
				StatusCode:        400,
				Headers:           map[string]string{},
				MultiValueHeaders: map[string][]string{},
				Body:              "Need body in 'POST' request",
				IsBase64Encoded:   false,
			}, nil
		}
		compiledMarkdown = compileMarkdown(request.Body)
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: compiledMarkdown}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        400,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              "Must either be GET or POST. Need 'text' query param in GET or body in POST",
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
