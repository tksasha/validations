package validations

import (
	"errors"
)

func ValidatePresenceOf(errs Errors, params ...string) error {
	err := errors.New("use: ValidatesPresenceOf(validations.Errors, attribue, value, message(optional))")

	var attribute, value, message string

	message = BLANK

	if len(params) > 0 {
		attribute = params[0]
	}

	if len(params) > 1 {
		value = params[1]
	}

	if len(params) > 2 {
		message = params[2]
	}

	if len(params) > 3 {
		return err
	}

	if value == "" {
		errs.Add(attribute, message)
	}

	return nil
}
