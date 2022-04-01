package errors

import (
	"net/http"
	"testing"
)

func TestNewInternalServerError_WhenError_ShouldMatchStringAndStatusCode(t *testing.T) {
	err := NewInternalServerError("can't process this request")
	expectedErr := "failed to process the request: can't process this request"
	if err.Message != expectedErr {
		t.Errorf("Expected %s; got %s", expectedErr, err.Message)
	}
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected %s; got %s", expectedErr, err.Message)
	}
}

func TestNewNilError_WhenError_ShouldMatchString(t *testing.T) {
	err := NewNilError("file")
	expectedErr := "file can not be nil"
	if err.Error() != expectedErr {
		t.Errorf("Expected %s; got %s", expectedErr, err.Error())
	}
}
