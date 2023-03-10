package main

import (
	"github.com/amlodzianowski/s3-cache-header-setter-go/header_setter"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(header_setter.HandleRequest)
}
