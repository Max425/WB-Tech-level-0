package constants

import "errors"

var (
	AlreadyExistsError = errors.New("unique constraint error")
)
