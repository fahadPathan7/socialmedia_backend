package errorcodes

import "errors"

var (
	// ErrNotFound error
	ErrNotFound = errors.New("not found")
	// ErrConnectionTimeout error
	ErrConnectionTimeout = errors.New("connection timeout")
	// ErrInvalidRequest error
	ErrInvalidRequest = errors.New("invalid request")
	// ErrUnauthorized error
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden error
	ErrForbidden = errors.New("forbidden")
	// ErrBadRequest error
	ErrBadRequest = errors.New("badrequest")
	// ErrPreconditionFailed error
	ErrPreconditionFailed = errors.New("precondition failed")
	// ErrPaymentFailed err type
	ErrPaymentFailed = errors.New("payment failed")
	// ErrCouldNotPerform err type
	ErrCouldNotPerform = errors.New("could not perform")
	// ErrAlreadyExists error def
	ErrAlreadyExists = errors.New("already exists")
)
