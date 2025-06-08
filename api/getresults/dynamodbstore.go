package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBStore struct {
	svc       *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBStore() *DynamoDBStore {
	sess := session.Must(session.NewSession())
	return &DynamoDBStore{
		svc:       dynamodb.New(sess),
		tableName: aws.String(os.Getenv("TableName")),
	}
}

func (d *DynamoDBStore) GetDepartmentData(waitGroup *sync.WaitGroup, target *map[string][]float32, dept string, errChan chan<- error) {
	defer waitGroup.Done()

	result := []float32{0, 0, 0, 0}

	scanInput := dynamodb.ScanInput{
		TableName: d.tableName,
		ScanFilter: map[string]*dynamodb.Condition{
			"sk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(fmt.Sprintf("DEPARTMENT#%s", dept)),
					},
				},
			},
		},
	}

	scanResult, err := d.svc.Scan(&scanInput)
	if err != nil {
		fmt.Printf("Failed to fetch items: %v\n", err)
		errChan <- fmt.Errorf("scan failed for dept %s: %w", dept, err)
		return
	}

	var responses []SurveyResponse
	err = dynamodbattribute.UnmarshalListOfMaps(scanResult.Items, &responses)
	if err != nil {
		errChan <- fmt.Errorf("unmarshal failed for dept %s: %w", dept, err)
		return
	}

	for _, response := range responses {
		for i := 0; i < 4; i++ {
			result[i] += float32(response.Results[i])
		}
	}

	if len(responses) > 0 {
		responseCount := float32(len(responses))
		for i := 0; i < 4; i++ {
			result[i] /= responseCount
		}
	}

	(*target)[dept] = result
	errChan <- nil
}
