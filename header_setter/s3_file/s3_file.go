package s3_file

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3File struct {
	BucketName string
	FileKey    string
	S3Service  *s3.Client
}

func New(BucketName string, FileKey string, S3Service *s3.Client) s3File {
	e := s3File{BucketName, FileKey, S3Service}
	return e
}

func (e s3File) HeadObject() (*s3.HeadObjectOutput, error) {
	headInput := &s3.HeadObjectInput{
		Bucket: &e.BucketName,
		Key:    &e.FileKey,
	}
	return e.S3Service.HeadObject(context.TODO(), headInput)
}

func (e s3File) ConfigureCacheControl(headObjectRes *s3.HeadObjectOutput) (*s3.CopyObjectOutput, error) {
	log.Printf("Configuring cache control on file: %v/%v", e.BucketName, e.FileKey)
	cacheControl := "max-age=31536000"
	copySource := fmt.Sprintf("/%v/%v", e.BucketName, e.FileKey)
	copyInput := &s3.CopyObjectInput{
		Bucket:            &e.BucketName,
		CopySource:        &copySource,
		Key:               &e.FileKey,
		CacheControl:      &cacheControl,
		ContentType:       headObjectRes.ContentType,
		MetadataDirective: "REPLACE",
		Metadata:          headObjectRes.Metadata,
		TaggingDirective:  "COPY",
	}
	return e.S3Service.CopyObject(context.TODO(), copyInput)
}
