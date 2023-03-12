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
	s3svc, svcErr := e.getS3Service()
	if svcErr != nil {
		return fmt.Sprintf("Unable to create session: %v", svcErr.Error()), 500
	}
	file := s3_file.New(e.BucketName, e.FileKey, s3svc)

	headObjectRes, headObjectErr := file.HeadObject()
	if headObjectErr != nil {
		return fmt.Sprintf("Unable to retrieve object head: %v", headObjectErr.Error()), 403
	}
	if headObjectRes.CacheControl != nil {
		return fmt.Sprintf("File: %v/%v already has a cache control setting of %v", e.BucketName, e.FileKey, *headObjectRes.CacheControl), 200
	}
	copyRes, copyErr := file.ConfigureCacheControl(headObjectRes)
	if copyErr != nil {
		return fmt.Sprintf("Unable to configure cache control: %v", copyErr.Error()), 403
	}
	return fmt.Sprintf("Event processed: %v", copyRes.CopyObjectResult), 201
}

func (e s3Event) getS3Service() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(e.AwsRegion)},
	)
	return s3.New(sess), err
}
