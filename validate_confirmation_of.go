package validations

import (
	"errors"
)

func ValidateConfirmationOf(errs Errors, params ...string) error {
	err := errors.New("use: ValidateConfirmationOf(attribute, value, confirmation, message(optional))")

	var attribute, value, confirmation, message string

	if len(params) > 0 {
		attribute = params[0]
	}

	if len(params) > 1 {
		value = params[1]
	}

	if len(params) > 2 {
		confirmation = params[2]
	}

	message = "doesn't match " + attribute + "_confirmation"

	if len(params) > 3 {
		message = params[3]
	}

	if len(params) > 4 {
		return err
	}

	if value != confirmation {
		errs.Add(attribute, message)
	}

	return nil
}
