package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type requestParameters struct {
	Email      string `json:"email"`
	Department string `json:"department"`
	Results    []int  `json:"results"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Got request: %s\n", request.Body)

	var requestData requestParameters
	if err := json.Unmarshal([]byte(request.Body), &requestData); err != nil {
		fmt.Printf("Failed to parse body: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	if requestData.Email == "" || requestData.Department == "" || len(requestData.Results) != 4 {
		fmt.Printf("Missing data from client: %v\n", requestData)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		StatusCode: 204,
		Headers:    headers,
	}, nil
}
