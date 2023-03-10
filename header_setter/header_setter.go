package header_setter

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}
func HandleRequest(_ context.Context, request events.S3Event) (events.APIGatewayProxyResponse, error) {
	record := request.Records[0]
	s3_bucket := record.S3.Bucket.Name
	s3_key := record.S3.Object.Key
	response := fmt.Sprintf("Received event to process following file: %v/%v", s3_bucket, s3_key)
	return GenerateResponse(response, 200), nil
}
