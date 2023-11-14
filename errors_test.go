package validations

import (
	"testing"
)

func TestAdd(t *testing.T) {
	errs := NewErrors()

	errs.Add("name", "can't be blank")

	subject := errs.Get("name")[0]

	expected := "can't be blank"

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}
}

func TestGet(t *testing.T) {
	errs := NewErrors()

	errs.Add("name", "can't be blank")

	subject := errs.Get("name")[0]

	expected := "can't be blank"

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}
}

func TestEmpty(t *testing.T) {
	errs := NewErrors()

	var subject, expected bool

	subject = errs.Empty()

	expected = true

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}

	errs.Add("name", "can't be blank")

	subject = errs.Empty()

	expected = false

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}
}

func TestPresent(t *testing.T) {
	errs := NewErrors()

	var subject, expected bool

	subject = errs.Present()

	expected = false

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}

	errs.Add("name", "can't be blank")

	subject = errs.Present()

	expected = true

	if subject != expected {
		t.Errorf(MESSAGE, subject, expected)
	}
}
