package validations

import (
	"testing"
)

func TestCantBeBlankMessage(t *testing.T) {
	subject := BLANK

	expected := "can't be blank"

	if subject != expected {
		t.Errorf(M, subject, expected)
	}
}
