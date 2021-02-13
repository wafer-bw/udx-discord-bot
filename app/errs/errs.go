package errs

import "errors"

// ErrUnauthorized - the request signature is invalid
var ErrUnauthorized = errors.New("unauthorized")

// ErrInvalidInteractionType - the request interaction type is invalid
var ErrInvalidInteractionType = errors.New("invalid interaction type")

// ErrNotImplemented - whatever was requested is not implemented yet
var ErrNotImplemented = errors.New("not implemented")

// ErrAlreadyExists - returned when attempting to create a command which already exists
var ErrAlreadyExists = errors.New("already exists")
