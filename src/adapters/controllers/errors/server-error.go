package errors

import "errors"

func ServerError(reason string) error {
	return errors.New(reason)
}