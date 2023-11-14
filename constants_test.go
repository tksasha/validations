package validations

import (
	"testing"
)

func TestCantBeBlankMessage(t *testing.T) {
	subject := CANT_BE_BLANK

	expected := "can't be blank"

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}
}
