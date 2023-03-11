package s3_file

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

type s3File struct {
	BucketName string
	FileKey    string
	S3Service  *s3.S3
	objectHead *s3.HeadObjectOutput
}

func New(BucketName string, FileKey string, S3Service *s3.S3) s3File {
	e := s3File{BucketName, FileKey, S3Service, nil}
	return e
}

func (e s3File) CacheControlConfigured() (bool, error) {
	res, err := e.headObject()
	if err != nil {
		return false, err
	}
	e.objectHead = res
	log.Printf("Cache control setting %v", res.CacheControl)
	return res.CacheControl != nil, err
}

func (e s3File) ConfigureCacheControl() {
	log.Printf("Configuring cache control on file: %v/%v", e.BucketName, e.FileKey)
}

func (e s3File) headObject() (*s3.HeadObjectOutput, error) {
	headInput := &s3.HeadObjectInput{
		Bucket: &e.BucketName,
		Key:    &e.FileKey,
	}
	return e.S3Service.HeadObject(headInput)
}
