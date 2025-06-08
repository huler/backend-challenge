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

type requestResponse struct {
	TotalCount int `json:"totalCount"`
}

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
	fmt.Printf("Got request: %+v\n", request)
	fmt.Printf("Got request: %s\n", request.Body)

	defaultHeaders := make(map[string]string)
	defaultHeaders["Access-Control-Allow-Origin"] = "*"

	var requestData requestParameters
	if err := json.Unmarshal([]byte(request.Body), &requestData); err != nil {
		fmt.Printf("Failed to parse body: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, err
	}

	if requestData.Email == "" || requestData.Department == "" || len(requestData.Results) != 4 {
		fmt.Printf("Missing data from client: %v\n", requestData)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}

	// Replaced `StoreSurveyResult` with `AddDepartmentSurveyResponse`
	err := store.AddDepartmentSurveyResponse(requestData)
	if err != nil {
		fmt.Printf("Failed to store survey result: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}

	// Replaced `BumpResponseCount` with `UpdateResponseCount`
	err = store.UpdateResponseCount()
	if err != nil {
		fmt.Printf("Failed to bump response count: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}

	totalCount, err := store.GetTotalResponseCount()
	if err != nil {
		fmt.Printf("Failed to fetch total response count: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}

	body, err := json.Marshal(requestResponse{TotalCount: totalCount})
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
