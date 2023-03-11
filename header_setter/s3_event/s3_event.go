package s3_event

import (
	"fmt"

	"github.com/amlodzianowski/s3-cache-header-setter-go/header_setter/s3_file"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Event struct {
	BucketName string
	FileKey    string
	AwsRegion  string
}

func New(BucketName string, FileKey string, AwsRegion string) s3Event {
	e := s3Event{BucketName, FileKey, AwsRegion}
	return e
}

func (e s3Event) ProcessEvent() (string, int) {
	s3svc, err := e.getS3Service()
	if err != nil {
		return fmt.Sprintf("Unable to create session: %v", err.Error()), 500
	}
	file := s3_file.New(e.BucketName, e.FileKey, s3svc)

	res, err := file.CacheControlConfigured()
	if err != nil {
		return fmt.Sprintf("Unable to retrieve cache control setting: %v", err.Error()), 403
	}
	if res {
		return fmt.Sprintf("File: %v/%v already has a cache control setting", e.BucketName, e.FileKey), 200
	}
	return fmt.Sprintf("Event processed: %v", res), 201
}

func (e s3Event) getS3Service() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(e.AwsRegion)},
	)
	return s3.New(sess), err
}
