package header_setter

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/aws/aws-lambda-go/events"
)

func TestGenerateResponse(t *testing.T) {

	got := GenerateResponse("Request successful", 200)
	want := events.APIGatewayProxyResponse{Body: "Request successful", StatusCode: 200}

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestHandleRequest(t *testing.T) {
	filename := "/../fixtures/s3_event.json"

	pwd, _ := os.Getwd()

	file, err := os.ReadFile(pwd + filename)
	if err != nil {
		t.Errorf("failed to load json file: %s, error: %v", filename, err)
		return
	}

	data := events.S3Event{}
	if err := json.Unmarshal(file, &data); err != nil {
		t.Errorf("failed to unmarshal json file, error: %v", err)
		return
	}

	// this actually tries to get AWS credentials and call the API, need to mock

	got, _ := HandleRequest(context.TODO(), data)
	want := 403

	if !cmp.Equal(got.StatusCode, want) {
		t.Errorf("got %v, wanted %v", got.StatusCode, want)
	}
}
