package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SurveyResponse struct {
	Pk      string `json:"pk"`
	Sk      string `json:"sk"`
	Results []int  `json:"results"`
}

var store SurveyStore

func main() {
	store = NewDynamoDBStore()
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Got request: %s\n", request.Body)

	defaultHeaders := make(map[string]string)
	defaultHeaders["Access-Control-Allow-Origin"] = "*"

	data := fetchData()

	body, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("Failed to marshall json: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    defaultHeaders,
		Body:       string(body),
	}, nil
}

func fetchData() map[string][]float32 {
	result := map[string][]float32{}

	departments := []string{
		"SAAS development",
		"Bespoke development",
		"Creative",
		"HR",
		"QA",
		"Project Management",
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(departments))

	for _, dept := range departments {
		go store.GetDepartmentData(&waitGroup, &result, dept)
	}

	waitGroup.Wait()

	return result
}
