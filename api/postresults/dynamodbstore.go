package main

import (
	"fmt"
	"os"

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

func (d *DynamoDBStore) AddDepartmentSurveyResponseWithCount(requestData requestParameters) error {
	surveyData := SurveyResponse{
		Pk:      fmt.Sprintf("RESPONSE#%s", requestData.Email),
		Sk:      fmt.Sprintf("DEPARTMENT#%s", requestData.Department),
		Results: requestData.Results,
	}

	toStore, err := dynamodbattribute.MarshalMap(surveyData)
	if err != nil {
		fmt.Printf("Failed to marshall manifest data: %v\n", err)
		return err
	}

	transactInput := &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					TableName: d.tableName,
					Item:      toStore,
				},
			},
			{
				Update: &dynamodb.Update{
					TableName: d.tableName,
					Key: map[string]*dynamodb.AttributeValue{
						"pk": {S: aws.String("RESPONSE-COUNT")},
						"sk": {S: aws.String("TOTAL")},
					},
					ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
						":inc":   {N: aws.String("1")},
						":start": {N: aws.String("0")},
					},
					UpdateExpression: aws.String("SET responses = if_not_exists(responses, :start) + :inc"),
				},
			},
		},
	}

	_, err = d.svc.TransactWriteItems(transactInput)
	if err != nil {
		return fmt.Errorf("failed to perform transaction: %w", err)
	}

	return nil
}

func (d *DynamoDBStore) AddDepartmentSurveyResponse(requestData requestParameters) error {
	surveyData := SurveyResponse{
		Pk:      fmt.Sprintf("RESPONSE#%s", requestData.Email),
		Sk:      fmt.Sprintf("DEPARTMENT#%s", requestData.Department),
		Results: requestData.Results,
	}

	toStore, err := dynamodbattribute.MarshalMap(surveyData)
	if err != nil {
		fmt.Printf("Failed to marshall manifest data: %v\n", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: d.tableName,
		Item:      toStore,
	}

	_, err = d.svc.PutItem(input)
	if err != nil {
		fmt.Printf("Failed to write to dynamo: %v\n", err)
		return err
	}

	return nil
}

func (d *DynamoDBStore) UpdateResponseCount() error {
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: d.tableName,
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

	_, err := d.svc.UpdateItem(updateItemInput)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDBStore) GetTotalResponseCount() (int, error) {
	var queryInput = &dynamodb.GetItemInput{
		TableName: d.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("RESPONSE-COUNT"),
			},
			"sk": {
				S: aws.String("TOTAL"),
			},
		},
	}

	data, err := d.svc.GetItem(queryInput)
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
		return 0, err
	}

	return item.Responses, nil
}
