package error

import (
	"testing"
)

func TestError_WhenError_Should_MatchString(t *testing.T) {
	expectedError := "mock error"
	err := NewError(expectedError)
	if err.Error() != expectedError {
		t.Errorf("Expected %s; got %s", expectedError, err.Error())
	}
}
