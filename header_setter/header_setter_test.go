package header_setter

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/aws/aws-lambda-go/events"
)

func TestGenerateResponse(t *testing.T) {

	got := GenerateResponse("Request successful", 200)
	want := events.APIGatewayProxyResponse{Body: "Request successful", StatusCode: 200}

	// if got != want {
	// 	t.Errorf("got %q, wanted %q", got, want)
	// }

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
