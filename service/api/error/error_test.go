package error

import (
	"testing"
)

func TestNewInternalServerError_WhenError_ShouldMatchStringAndStatusCode(t *testing.T) {
	err := NewInternalServerError("request can't be processed")
	expectedErr := "internal server error: request can't be processed"
	if err.Error() != expectedErr {
		t.Errorf("Expected %s; got %s", expectedErr, err.Error())

	}
}
