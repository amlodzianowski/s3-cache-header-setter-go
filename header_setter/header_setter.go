package header_setter

import (
	"context"
	"log"

	"github.com/amlodzianowski/s3-cache-header-setter-go/header_setter/s3_event"
	"github.com/aws/aws-lambda-go/events"
)

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}
func HandleRequest(_ context.Context, request events.S3Event) (events.APIGatewayProxyResponse, error) {
	record := request.Records[0]
	event := s3_event.New(record.S3.Bucket.Name, record.S3.Object.Key, record.AWSRegion)
	log.Printf("Received event to process following file: %v/%v in region %v", event.BucketName, event.FileKey, event.AwsRegion)
	response, code := event.ProcessEvent()
	log.Print(response)
	return GenerateResponse(response, code), nil
}
