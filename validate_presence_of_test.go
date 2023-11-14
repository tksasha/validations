package validations

import (
	"testing"
)

func TestValidatePresenceOf(t *testing.T) {
	var errs Errors

	t.Run("when `name` is not present", func(t *testing.T) {
		var subject, expected string

		errs = NewErrors()

		name := ""

		ValidatePresenceOf(errs, "name", name)

		subject = errs.Get("name")[0]

		expected = "can't be blank"

		if subject != expected {
			t.Errorf(M, subject, expected)
		}
	})

	t.Run("when `name` is present", func(t *testing.T) {
		var subject, expected bool

		errs = NewErrors()

		name := "John McClane"

		ValidatePresenceOf(errs, "name", name)

		subject = errs.IsEmpty()

		expected = true

		if subject != expected {
			t.Errorf(M, subject, expected)
		}
	})
}
