package errs

import "errors"

var (
	ErrInvalidID           = errors.New("invalid id")
	ErrInvalidBodyRequest  = errors.New("invalid request body")
	ErrInternalServerError = errors.New("an error occurred internally")
)
