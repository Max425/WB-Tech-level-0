package constants

import "errors"

var (
	InvalidInputBodyError = errors.New("invalid input body")
	AlreadyExistsError    = errors.New("unique constraint error")
)
