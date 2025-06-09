package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type SurveyResponse struct {
	Pk      string `json:"pk"`
	Sk      string `json:"sk"`
	Results []int  `json:"results"`
}

func main() {
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
		go fetchDepartmentData(&waitGroup, &result, dept)
	}

	waitGroup.Wait()

	return result
}

type DepartmentStats struct {
	TotalScore int `json:"totalScore"`
	Responses  int `json:"responses"`
}

func fetchDepartmentData(waitGroup *sync.WaitGroup, target *map[string][]float32, dept string) {
	defer waitGroup.Done()

	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TableName")),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(fmt.Sprintf("STATS#DEPARTMENT#%s", dept)),
			},
			"sk": {
				S: aws.String("TOTALS"),
			},
		},
	}

	svc := dynamodb.New(session.New())
	data, err := svc.GetItem(getInput)
	if err != nil {
		fmt.Printf("Failed to fetch stats for %s: %v\n", dept, err)
		(*target)[dept] = []float32{0, 0, 0, 0}
		return
	}

	var stats DepartmentStats
	if err := dynamodbattribute.UnmarshalMap(data.Item, &stats); err != nil {
		fmt.Printf("Failed to unmarshal stats for %s: %v\n", dept, err)
		(*target)[dept] = []float32{0, 0, 0, 0}
		return
	}

	avg := float32(0)
	if stats.Responses > 0 {
		avg = float32(stats.TotalScore) / float32(stats.Responses)
	}

	(*target)[dept] = []float32{avg, avg, avg, avg}
}
