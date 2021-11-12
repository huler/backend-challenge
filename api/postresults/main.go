package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func main() {
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

	if err := storeSurveyResult(requestData); err != nil {
		fmt.Printf("Failed to store survey result: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}
	if err := bumpResponseCount(); err != nil {
		fmt.Printf("Failed to bump response count: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}
	totalCount, err := getTotalResponseCount()
	if err != nil {
		fmt.Printf("Failed to fetch total response count: %v\n", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    defaultHeaders,
		}, nil
	}

	response := requestResponse{
		TotalCount: totalCount,
	}
	body, err := json.Marshal(&response)
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
func storeSurveyResult(requestData requestParameters) error {
	surveyData := SurveyResponse{
		Pk:      fmt.Sprintf("RESPONSE#%s", requestData.Email),
		Sk:      fmt.Sprintf("DEPARTMENT#%s", requestData.Department),
		Results: requestData.Results,
	}

	toStore, marshalErr := dynamodbattribute.MarshalMap(surveyData)

	if marshalErr != nil {
		fmt.Printf("Failed to marshall manifest data: %v\n", marshalErr)
		return marshalErr
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TableName")),
		Item:      toStore,
	}

	svc := dynamodb.New(session.New())
	if _, err := svc.PutItem(input); err != nil {
		fmt.Printf("Failed to write to dynamo: %v\n", err)
		return err
	}

	return nil
}
func bumpResponseCount() error {
	var updateItemInput = &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("RESPONSE-COUNT"),
			},
			"sk": {
				S: aws.String("TOTAL"),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				N: aws.String("1"),
			},
			":start": {
				N: aws.String("0"),
			},
		},
		UpdateExpression: aws.String("SET responses = if_not_exists(responses, :start) + :inc"),
	}

	svc := dynamodb.New(session.New())
	_, err := svc.UpdateItem(updateItemInput)

	return err
}
func getTotalResponseCount() (int, error) {
	var queryInput = &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("RESPONSE-COUNT"),
			},
			"sk": {
				S: aws.String("TOTAL"),
			},
		},
	}
	svc := dynamodb.New(session.New())
	data, err := svc.GetItem(queryInput)
	if err != nil {
		return 0, err
	}

	type totalCount struct {
		Responses int `json:"responses"`
	}

	item := totalCount{}
	err = dynamodbattribute.UnmarshalMap(data.Item, &item)
	if err != nil {
		fmt.Printf("Failed to unmarshal item: %v\n", err)
	}

	return item.Responses, nil
}
