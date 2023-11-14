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
		t.Errorf(M, subject, expected)
	}
}

func TestGet(t *testing.T) {
	errs := NewErrors()

	errs.Add("name", "can't be blank")

	subject := errs.Get("name")[0]

	expected := "can't be blank"

	if subject != expected {
		t.Errorf(M, subject, expected)
	}
}

func TestIsEmpty(t *testing.T) {
	errs := NewErrors()

	var subject, expected bool

	subject = errs.IsEmpty()

	expected = true

	if subject != expected {
		t.Errorf(M, subject, expected)
	}

	errs.Add("name", "can't be blank")

	subject = errs.IsEmpty()

	expected = false

	if subject != expected {
		t.Errorf(M, subject, expected)
	}
}

func TestIsPresent(t *testing.T) {
	errs := NewErrors()

	var subject, expected bool

	subject = errs.IsPresent()

	expected = false

	if subject != expected {
		t.Errorf(M, subject, expected)
	}

	errs.Add("name", "can't be blank")

	subject = errs.IsPresent()

	expected = true

	if subject != expected {
		t.Errorf(M, subject, expected)
	}
}
